package wasm

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io/fs"
	"testing"
	"time"

	"github.com/TecharoHQ/anubis/web"
)

func abiTest(t testing.TB, fname string, difficulty uint32) {
	fin, err := web.Static.Open("static/wasm/" + fname)
	if err != nil {
		t.Fatal(err)
	}
	defer fin.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	t.Cleanup(cancel)

	runner, err := NewRunner(ctx, fname, fin)
	if err != nil {
		t.Fatal(err)
	}

	h := sha256.New()
	fmt.Fprint(h, t.Name())
	data := h.Sum(nil)

	nonce, hash, mod, err := runner.run(ctx, data, difficulty, 0, 1)
	if err != nil {
		t.Fatal(err)
	}

	if err := runner.writeVerification(ctx, mod, hash); err != nil {
		t.Fatalf("can't write verification: %v", err)
	}

	ok, err := runner.anubisValidate(ctx, mod, nonce, difficulty)
	if err != nil {
		t.Fatalf("can't run validation: %v", err)
	}

	if !ok {
		t.Error("validation failed")
	}

	t.Logf("used %d pages of wasm memory (%d bytes)", mod.Memory().Size()/63356, mod.Memory().Size())
}

func TestAlgos(t *testing.T) {
	fnames, err := fs.ReadDir(web.Static, "static/wasm")
	if err != nil {
		t.Fatal(err)
	}

	for _, fname := range fnames {
		fname := fname
		t.Run(fname.Name(), func(t *testing.T) {
			abiTest(t, fname.Name(), 4)
		})
	}
}

func bench(b *testing.B, fname string, difficulties []uint32) {
	b.Helper()

	fin, err := web.Static.Open("static/wasm/" + fname)
	if err != nil {
		b.Fatal(err)
	}
	defer fin.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	b.Cleanup(cancel)

	runner, err := NewRunner(ctx, fname, fin)
	if err != nil {
		b.Fatal(err)
	}

	h := sha256.New()
	fmt.Fprint(h, "This is an example value that exists only to test the system.")
	data := h.Sum(nil)

	_, _, mod, err := runner.run(ctx, data, 0, 0, 1)
	if err != nil {
		b.Fatal(err)
	}

	for _, difficulty := range difficulties {
		b.Run(fmt.Sprintf("difficulty/%d", difficulty), func(b *testing.B) {
			for b.Loop() {
				difficulty := difficulty
				_, err := runner.anubisWork(ctx, mod, difficulty, 0, 1)
				if err != nil {
					b.Fatalf("can't do test work run: %v", err)
				}
			}
		})
	}
}

func BenchmarkSHA256(b *testing.B) {
	bench(b, "sha256.wasm", []uint32{4, 6, 8, 10, 12, 14, 16})
}

func BenchmarkArgon2ID(b *testing.B) {
	bench(b, "argon2id.wasm", []uint32{4, 6, 8})
}

func BenchmarkValidate(b *testing.B) {
	fnames, err := fs.ReadDir(web.Static, "static/wasm")
	if err != nil {
		b.Fatal(err)
	}

	h := sha256.New()
	fmt.Fprint(h, "This is an example value that exists only to test the system.")
	data := h.Sum(nil)

	for _, fname := range fnames {
		fname := fname.Name()

		difficulty := uint32(1)

		switch fname {
		case "sha256.wasm":
			difficulty = 16
		}

		b.Run(fname, func(b *testing.B) {
			fin, err := web.Static.Open("static/wasm/" + fname)
			if err != nil {
				b.Fatal(err)
			}
			defer fin.Close()
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			b.Cleanup(cancel)

			runner, err := NewRunner(ctx, fname, fin)
			if err != nil {
				b.Fatal(err)
			}

			nonce, hash, mod, err := runner.run(ctx, data, difficulty, 0, 1)
			if err != nil {
				b.Fatal(err)
			}

			if err := runner.writeVerification(ctx, mod, hash); err != nil {
				b.Fatalf("can't write verification: %v", err)
			}

			for b.Loop() {
				_, err := runner.anubisValidate(ctx, mod, nonce, difficulty)
				if err != nil {
					b.Fatalf("can't run validation: %v", err)
				}
			}
		})
	}
}
