export const en = {
  template: "Configuration template (new profile only)",
  chooseTemplate: "Choose a template…",
  general: "Z.ai GLM-5.3 · General API (recommended)",
  coding: "Z.ai GLM-5.3 · Coding Plan (reference)",
  introduction:
    "Optional: prefill connection fields, then enter a profile name and your API key. Your key, proxy, and name are kept. Selection does not save or activate a profile. Save when ready, then activate separately. The model remains editable.",
  generalHelp:
    "Uses the General API endpoint and API billing, separate from Coding Plan subscription quota. Confirm model access and billing for your Z.ai account.",
  codingHelp:
    "Reference only: Coding Plan uses a separate endpoint and subscription quota. Z.ai restricts it to officially supported tools; ScopeWeaver is not listed. Obtain Z.ai authorization before using this endpoint with ScopeWeaver.",
  modelHelp:
    "GLM-5.3 is text-only with a 1M-token context. Reasoning must stay enabled; supported effort levels are low, high, and max. This template selects max and leaves the output limit at 0 (provider default). Review these settings if you change the model.",
  officialModel: "GLM-5.3 model guide",
  officialPolicy: "Coding Plan usage policy",
};
export const ko: Record<keyof typeof en, string> = {
  template: "설정 템플릿(새 프로필 전용)",
  chooseTemplate: "템플릿 선택…",
  general: "Z.ai GLM-5.3 · 일반 API(권장)",
  coding: "Z.ai GLM-5.3 · Coding Plan(참고용)",
  introduction:
    "선택하면 연결 설정을 미리 입력합니다. 프로필 이름과 본인의 API 키를 입력하세요. 기존에 입력한 키, 프록시, 이름은 유지하며 선택만으로 저장하거나 활성화하지 않습니다. 준비되면 저장하고 별도로 활성화하세요. 모델 이름은 수정할 수 있습니다.",
  generalHelp:
    "일반 API 엔드포인트와 API 요금을 사용하며 Coding Plan 구독 할당량과 별개입니다. Z.ai 계정의 모델 접근 권한과 과금 조건을 확인하세요.",
  codingHelp:
    "참고용: Coding Plan은 별도 엔드포인트와 구독 할당량을 사용합니다. Z.ai가 공식 지원 도구로 사용을 제한하며 ScopeWeaver는 목록에 없습니다. ScopeWeaver에서 이 엔드포인트를 사용하려면 먼저 Z.ai의 허가를 받으세요.",
  modelHelp:
    "GLM-5.3은 텍스트 전용이며 문맥 한도는 100만 토큰입니다. 추론은 항상 켜야 하며 지원 강도는 low, high, max입니다. 템플릿은 max를 선택하고 출력 한도는 0(제공자 기본값)으로 둡니다. 모델을 바꾸면 이 설정도 확인하세요.",
  officialModel: "GLM-5.3 모델 안내",
  officialPolicy: "Coding Plan 이용 정책",
};
export const ja: Record<keyof typeof en, string> = {
  template: "設定テンプレート（新規プロファイルのみ）",
  chooseTemplate: "テンプレートを選択…",
  general: "Z.ai GLM-5.3 · 一般 API（推奨）",
  coding: "Z.ai GLM-5.3 · Coding Plan（参考）",
  introduction:
    "任意: 接続項目を事前入力したうえで、プロファイル名とご自身の API キーを入力してください。入力済みのキー・プロキシ・名前は保持されます。選択だけではプロファイルの保存も有効化も行われません。準備ができたら保存し、別途有効化してください。モデルは引き続き編集できます。",
  generalHelp:
    "一般 API のエンドポイントと API 課金を使用し、Coding Plan のサブスクリプション枠とは別です。Z.ai アカウントのモデル利用権と課金条件をご確認ください。",
  codingHelp:
    "参考のみ: Coding Plan は別のエンドポイントとサブスクリプション枠を使用します。Z.ai は公式サポート対象のツールに利用を限定しており、ScopeWeaver は一覧に含まれていません。ScopeWeaver でこのエンドポイントを使用する前に、Z.ai の許可を得てください。",
  modelHelp:
    "GLM-5.3 はテキスト専用で、コンテキストは 100 万トークンです。推論は常に有効にする必要があり、対応する強度は low・high・max です。このテンプレートは max を選択し、出力上限は 0（プロバイダー既定値）のままにします。モデルを変更した場合はこれらの設定を見直してください。",
  officialModel: "GLM-5.3 モデルガイド",
  officialPolicy: "Coding Plan 利用ポリシー",
};
