package locale

import (
	"strings"
	"testing"
)

// TestJapaneseFormatCompatible verifies that every catalog entry carrying a
// Japanese translation keeps the same printf argument signature as its English
// source, mirroring the En/Ko check in TestProductionCatalogCompleteAndFormatCompatible.
// A mismatch would crash or mis-render a localized message at runtime.
func TestJapaneseFormatCompatible(t *testing.T) {
	checked := 0
	for key, values := range catalog {
		ja, ok := values[Ja]
		if !ok {
			continue
		}
		en, has := values[En]
		if !has {
			continue
		}
		enSig := strings.Join(formatArguments(en), ",")
		jaSig := strings.Join(formatArguments(ja), ",")
		if enSig != jaSig {
			t.Errorf("format arguments differ for %q: English %s; Japanese %s", key, enSig, jaSig)
		}
		checked++
	}
	if checked < 200 {
		t.Fatalf("too few Japanese translations verified: %d (expected the main backend surface)", checked)
	}
	t.Logf("verified printf parity for %d Japanese entries", checked)
}
