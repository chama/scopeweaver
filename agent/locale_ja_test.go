package agent

import (
	"strings"
	"testing"

	"github.com/Autumn-27/artex/locale"
)

// TestBuiltinPromptJapaneseCoverage mirrors TestBuiltinPromptTranslationCoverage
// for Japanese: every built-in prompt must have a non-identity Japanese
// translation that preserves the Go-template variable and the tool-name
// contracts, so the Japanese-localized agents behave identically to English.
func TestBuiltinPromptJapaneseCoverage(t *testing.T) {
	for _, p := range []string{goalsDefaultTmpl, goalsScopeTail, plannerDefaultTmpl, workerDefaultTmpl, mainAgentDefaultTmpl, autoDefaultTmpl, pentestDefaultTmpl, DefaultAssistantPrompt, ReporterDefaultPrompt, RetesterDefaultPrompt} {
		ja, ok := locale.Lookup(locale.Ja, p)
		if !ok || ja == p || strings.TrimSpace(ja) == "" {
			t.Fatalf("missing Japanese prompt: %.60s", p)
		}
		for _, contract := range []string{"{{.Goal}}", "set_constraints", "add_task_scope", "report_finding", "update_finding_report", "record_finding_retest_result"} {
			if strings.Contains(p, contract) != strings.Contains(ja, contract) {
				t.Fatalf("Japanese translation lost contract %q for %.40s", contract, p)
			}
		}
	}
}

// TestOutputLanguageInstructionJapanese confirms the Japanese branch is wired
// and selects Japanese without affecting English/Korean.
func TestOutputLanguageInstructionJapanese(t *testing.T) {
	ja := outputLanguageInstruction(locale.Ja)
	if !strings.Contains(ja, "日本語") {
		t.Fatalf("Japanese output-language instruction not selected: %q", ja)
	}
	if en := outputLanguageInstruction(locale.En); !strings.Contains(en, "English") {
		t.Fatalf("English instruction regressed: %q", en)
	}
}
