package locale

// Japanese translations for core startup/config/LLM-pool/platform-control
// messages (see core.go). High-visibility: these appear in the startup banner
// and the System logs page.
func init() {
	RegisterJa("environment variable ARTEX_PG_DSN", "環境変数 ARTEX_PG_DSN")
	RegisterJa("configuration file ", "設定ファイル ")
	RegisterJa(" (database fields)", "（database フィールド）")
	RegisterJa("Database configuration not found: ARTEX_PG_DSN is unset and %s has no database configuration (dsn or host/user/dbname). Create the configuration file or set the environment variable and retry", "データベース設定が見つかりません: ARTEX_PG_DSN が未設定で、%s にも database 設定（dsn または host/user/dbname）がありません。設定ファイルを作成するか環境変数を設定して再試行してください")
	RegisterJa("Evidence body verification failed: %s", "証拠本文の検証に失敗しました: %s")
	RegisterJa("Incomplete body: expected %d bytes, read %d bytes", "本文が不完全です: 予期したバイト数 %d、読み取ったバイト数 %d")
	RegisterJa("Raw traffic body hash mismatch", "生トラフィック本文のハッシュが一致しません")
	RegisterJa("side must be request or response", "side は request または response である必要があります")
	RegisterJa("Archived evidence snapshot metadata hash mismatch", "アーカイブされた証拠スナップショットのメタデータのハッシュが一致しません")
	RegisterJa("  Autonomous AI penetration-testing system", "  自律型 AI 侵入テストシステム")
	RegisterJa("  Version %s  ·  %s/%s  ·  %s  ·  Listening on %s\n\n", "  バージョン %s  ·  %s/%s  ·  %s  ·  待ち受け %s\n\n")
	RegisterJa("[config] Configuration file: %s", "[config] 設定ファイル: %s")
	RegisterJa("[config] Configuration file: %s (not found; using ARTEX_PG_DSN only)", "[config] 設定ファイル: %s（見つかりません。ARTEX_PG_DSN のみを使用します）")
	RegisterJa("[config] Skill directory: %s", "[config] スキルディレクトリ: %s")
	RegisterJa("Work was cancelled; platform security controls blocked execution", "作業がキャンセルされました。プラットフォームのセキュリティ制御が実行をブロックしました")
	RegisterJa("Human approval was not granted (denied or timed out)", "人間による承認が得られませんでした（拒否またはタイムアウト）")
	RegisterJa("[ScopeWeaver platform control, not target defense] This call was blocked by the platform. ", "[ScopeWeaver のプラットフォーム制御であり、対象の防御ではありません] この呼び出しはプラットフォームによってブロックされました。 ")
	RegisterJa("Reason: ", "理由: ")
	RegisterJa(". This operation is prohibited.", "。この操作は禁止されています。")
	RegisterJa("[enrich] dnsx initialization failed; DNS resolution disabled: %v", "[enrich] dnsx の初期化に失敗したため、DNS 解決を無効にしました: %v")
	RegisterJa("[enrich] Queue full; dropping job kind=%d id=%d", "[enrich] キューが満杯のため、ジョブを破棄します kind=%d id=%d")
	RegisterJa("LLM failover: all profiles are unavailable", "LLM フェイルオーバー: すべてのプロファイルが利用できません")
	RegisterJa("[llmpool] Profile %q(%s) circuit opened: %s", "[llmpool] プロファイル %q(%s) のサーキットが作動しました: %s")
	RegisterJa("[llmpool] LLM failover: %q(%s) -> %q(%s), reason: %s", "[llmpool] LLM フェイルオーバー: %q(%s) -> %q(%s)、理由: %s")
	RegisterJa("[llmpool] Failover chain exhausted (all %d profiles failed), last error: %s", "[llmpool] フェイルオーバーチェーンを使い果たしました（%d 個のプロファイルすべてが失敗）、最後のエラー: %s")
}
