package partials

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestHeader(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = Header().Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	if doc.Find(`[data-testid="header"]`).Length() == 0 {
		t.Error("expected data-testid attribute on header")
	}
}

func TestFooter(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = Footer().Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	if doc.Find(`[data-testid="footer"]`).Length() == 0 {
		t.Error("expected data-testid attribute on footer")
	}
}

func TestBody(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = Body().Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	if doc.Find(`[data-testid="header"]`).Length() == 0 {
		t.Error("expected data-testid attribute on header")
	}

	if doc.Find(`[data-testid="main"]`).Length() == 0 {
		t.Error("expected data-testid attribute on main")
	}

	if doc.Find(`[data-testid="footer"]`).Length() == 0 {
		t.Error("expected data-testid attribute on footer")
	}
}
