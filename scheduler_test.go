package job_test

import (
	"context"
	"sync"
	"testing"

	"github.com/ne2blink/job-scheduler"
)

func BenchmarkQueueScheduler1(b *testing.B) {
	s := job.NewQueueScheduler(1)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkQueueScheduler2(b *testing.B) {
	s := job.NewQueueScheduler(2)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkQueueScheduler4(b *testing.B) {
	s := job.NewQueueScheduler(4)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkQueueScheduler8(b *testing.B) {
	s := job.NewQueueScheduler(8)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkQueueScheduler16(b *testing.B) {
	s := job.NewQueueScheduler(16)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkQueueScheduler32(b *testing.B) {
	s := job.NewQueueScheduler(32)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkQueueScheduler64(b *testing.B) {
	s := job.NewQueueScheduler(64)

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler1(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 1)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler2(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 2)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler4(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 4)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler8(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 8)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler16(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 16)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler32(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 32)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func BenchmarkChannelScheduler64(b *testing.B) {
	s := job.NewChannelScheduler(context.Background(), 64)
	s.Start()

	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		s.Add(func() {
			wg.Done()
		})
	}
	wg.Wait()
}
