// Assembles the per-namespace message modules into flat en/ko/ja catalogs.
// Keys are "<namespace>.<key>". Every namespace module types its Korean and
// Japanese tables as Record<keyof typeof en, string>, so a missing or extra
// entry is a compile error; i18n.test.mjs additionally checks placeholder/tag
// parity. Japanese tables are built as { ...en, ...overrides } so every key is
// always present, with English standing in until a Japanese string is authored.

import * as app from "./messages/app.ts";
import * as common from "./messages/common.ts";
import * as english from "./messages/english.ts";
import * as interfaceMessages from "./messages/interface.ts";
import * as providers from "./messages/providers.ts";
import * as settings from "./messages/settings.ts";

const namespaces = {
  providers,
  settings,
  english,
  interface: interfaceMessages,
  app,
  common,
};

type Namespaces = typeof namespaces;
type NamespaceName = keyof Namespaces & string;

export type MessageKey = {
  [N in NamespaceName]: `${N}.${keyof Namespaces[N]["en"] & string}`;
}[NamespaceName];

export type Catalog = Record<MessageKey, string>;

function flatten(locale: "en" | "ko" | "ja"): Catalog {
  const out: Record<string, string> = {};
  for (const [ns, mod] of Object.entries(namespaces)) {
    for (const [key, value] of Object.entries(mod[locale] as Record<string, string>)) out[`${ns}.${key}`] = value;
  }
  return out as Catalog;
}

export const catalogs = { en: flatten("en"), ko: flatten("ko"), ja: flatten("ja") } as const;

export const namespaceNames = Object.keys(namespaces);
