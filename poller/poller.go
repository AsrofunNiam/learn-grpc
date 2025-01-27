package poller

import (
	"context"
	"log"
	"time"
)

// Error untuk konfigurasi tidak valid
var ErrInvalidConfiguration = context.DeadlineExceeded

// Worker interface mendefinisikan pekerjaan yang akan diproses oleh poller
type Worker interface {
	Name() string             // Nama worker
	Work(ctx context.Context) // Pekerjaan yang akan diproses
	OnError(err error)        // Callback untuk error
	Interval() time.Duration  // Interval polling
}

// Poller adalah tipe utama untuk menjalankan worker secara periodik
type Poller struct {
	name       string
	work       func(ctx context.Context)
	onError    func(err error)
	noInit     bool
	interval   time.Duration
	cancelFunc context.CancelFunc
}

// Option untuk konfigurasi tambahan poller
type Option func(*Poller)

// WithNoInitialization adalah opsi untuk menghilangkan inisialisasi awal
func WithNoInitialization() Option {
	return func(p *Poller) {
		p.noInit = true
	}
}

// New membuat instansi poller baru
func New(name string, work func(ctx context.Context), onError func(err error), opts ...Option) (*Poller, error) {
	if name == "" || work == nil || onError == nil {
		return nil, ErrInvalidConfiguration
	}

	p := &Poller{
		name:    name,
		work:    work,
		onError: onError,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p, nil
}

// Start menjalankan worker secara periodik
func (p *Poller) Start(ctx context.Context, interval time.Duration) {
	p.interval = interval
	ctx, cancel := context.WithCancel(ctx)
	p.cancelFunc = cancel

	go func() {
		ticker := time.NewTicker(p.interval)
		defer ticker.Stop()

		if !p.noInit {
			// Jalankan pekerjaan pertama kali jika opsi NoInitialization tidak diaktifkan
			if err := p.run(ctx); err != nil {
				p.onError(err)
			}
		}

		for {
			select {
			case <-ticker.C:
				if err := p.run(ctx); err != nil {
					p.onError(err)
				}
			case <-ctx.Done():
				log.Printf("Poller [%s] stopped", p.name)
				return
			}
		}
	}()
}

// run adalah helper untuk menjalankan pekerjaan worker
func (p *Poller) run(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	p.work(ctx)
	return nil
}

// Stop menghentikan poller
func (p *Poller) Stop() {
	if p.cancelFunc != nil {
		p.cancelFunc()
	}
	log.Printf("Poller [%s] has been stopped", p.name)
}
