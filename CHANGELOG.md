# Changelog

## [1.5.0](https://github.com/run-llama/llama-parse-go/compare/v1.4.0...v1.5.0) (2026-08-12)

> [!WARNING]
> **Breaking change:** `FileService.Get` is renamed to `FileService.Content`; it still returns `*PresignedURL`.
> `FileService.Get` now returns `*FileGetResponse` (the file resource), so existing `client.Files.Get(...)` calls fail to compile rather than misbehave silently. Replace them with `client.Files.Content(...)`.
> Released as a minor, not a major, to avoid forcing the module path to `github.com/run-llama/llama-parse-go/v2`.


### Features

* **batches:** webhooks for /api/v2/batches ([#23148](https://github.com/run-llama/llama-parse-go/issues/23148)) ([4a35d5d](https://github.com/run-llama/llama-parse-go/commit/4a35d5d95a92571e8f1db859cbe6883784142050))
* **classify:** accept webhook_configuration_ids on classify job create (LI-8138) ([#22943](https://github.com/run-llama/llama-parse-go/issues/22943)) ([e86755e](https://github.com/run-llama/llama-parse-go/commit/e86755e508a479c6be4dfc0384edc65b0482d530))
* **connector:** API + service layer for attaching a subscription to a directory ([#23502](https://github.com/run-llama/llama-parse-go/issues/23502)) ([c5b77e8](https://github.com/run-llama/llama-parse-go/commit/c5b77e81d77a069a93675864cff14b1c85e81e57))
* **connectors:** expose a directory's connector subscription, and resolve connected accounts ([#23831](https://github.com/run-llama/llama-parse-go/issues/23831)) ([60882f3](https://github.com/run-llama/llama-parse-go/commit/60882f3b0604b67d6f95b017e714030852e833ad))
* **extract:** accept webhook_configuration_ids on extract job create (LI-8138) ([#22907](https://github.com/run-llama/llama-parse-go/issues/22907)) ([b344911](https://github.com/run-llama/llama-parse-go/commit/b344911cd20dbafac08e16903ccda830c408888a))
* **extract:** opt-in spreadsheet mode for agentic_plus ([#22958](https://github.com/run-llama/llama-parse-go/issues/22958)) ([4cc2e32](https://github.com/run-llama/llama-parse-go/commit/4cc2e32c38528daf7d2153b56118973f9f8f4613))
* **extract:** pin turbo to a stable dated version; accept citations+confidence, reject only granular bboxes ([#22965](https://github.com/run-llama/llama-parse-go/issues/22965)) ([9e67e18](https://github.com/run-llama/llama-parse-go/commit/9e67e186abac557a9a074bf4d9d8080418ebe257))
* **extract:** reject parse_tier for parse-free tiers + pin turbo fallback to fast ([#22919](https://github.com/run-llama/llama-parse-go/issues/22919)) ([dd435a1](https://github.com/run-llama/llama-parse-go/commit/dd435a17357972965c33385186ec6f65342b60fe))
* **files:** rename files.get to files.content and restore files.retrieve ([96cd0ff](https://github.com/run-llama/llama-parse-go/commit/96cd0ff1de81cf09eb6fdbf0ef1995bcba83c993))
* **forms:** emit bboxes in v2 forms output ([#23974](https://github.com/run-llama/llama-parse-go/issues/23974)) ([6ad27a9](https://github.com/run-llama/llama-parse-go/commit/6ad27a93ecd40dddcbc6d6ca7a2db220c0227be0))
* **parse,extract:** add expand=usage returning credits billed per job ([#23709](https://github.com/run-llama/llama-parse-go/issues/23709)) ([c230b00](https://github.com/run-llama/llama-parse-go/commit/c230b009630d04f13ab21bc01f5ad8996ed477a3))
* **parse:** extract Word revision annotations ([#23152](https://github.com/run-llama/llama-parse-go/issues/23152)) ([e553748](https://github.com/run-llama/llama-parse-go/commit/e5537487cfa2bda274348e673f87d0dab4f085b2))
* **parse:** make the output.pdf artifact opt-in on Parse v2 (output_options.save_output_pdf) ([#23510](https://github.com/run-llama/llama-parse-go/issues/23510)) ([392f917](https://github.com/run-llama/llama-parse-go/commit/392f917b68502a67d7168b85d800893095e538b1))
* **split:** accept webhook_configuration_ids on split job create (LI-8138) ([#22940](https://github.com/run-llama/llama-parse-go/issues/22940)) ([22839b9](https://github.com/run-llama/llama-parse-go/commit/22839b9a0fe1d5ec577f30e724e59ebc5547e137))


### Bug Fixes

* **api:** name paginated + filter schemas instead of leaking generic parametrizations ([#23120](https://github.com/run-llama/llama-parse-go/issues/23120)) ([5e1fb23](https://github.com/run-llama/llama-parse-go/commit/5e1fb2349f19708c46730335a2b066db78e16612))
* **llamaparse:** retry qwen context-overflow 400s with a shrinking OCR anchor ([#23817](https://github.com/run-llama/llama-parse-go/issues/23817)) ([c93c062](https://github.com/run-llama/llama-parse-go/commit/c93c0629c2290f5b94e6874443d611277222c644))


### Chores

* **api:** regenerate OpenAPI specs for new agentic parse version ([#22763](https://github.com/run-llama/llama-parse-go/issues/22763)) ([ea16a62](https://github.com/run-llama/llama-parse-go/commit/ea16a621051b3abebfc50c1f106918f2a235b7e5))


### Documentation

* **parse:** shorten the images_to_save field description ([#23807](https://github.com/run-llama/llama-parse-go/issues/23807)) ([9733f82](https://github.com/run-llama/llama-parse-go/commit/9733f826e9d5a0d517dac9bec9f0e8e7245910b7))


### Refactors

* remove Depends(get_db) from permissions endpoints ([#22635](https://github.com/run-llama/llama-parse-go/issues/22635)) ([c917909](https://github.com/run-llama/llama-parse-go/commit/c9179090d4d78128e6d93d13b0974518a5cd18a9))

## [1.4.0](https://github.com/run-llama/llama-parse-go/compare/v1.3.0...v1.4.0) (2026-07-22)


### Features

* **extract:** report num_pages_billed on extract job usage ([#22323](https://github.com/run-llama/llama-parse-go/issues/22323)) ([673894d](https://github.com/run-llama/llama-parse-go/commit/673894d764b14d3eeeb437028ffa574f7f07974d))
* **sheets:** cost_effective/agentic tiers and per-region billing ([#22508](https://github.com/run-llama/llama-parse-go/issues/22508)) ([e156f54](https://github.com/run-llama/llama-parse-go/commit/e156f54b7f1d083873f3da7a24f3091e49e9e358))

## [1.3.0](https://github.com/run-llama/llama-parse-go/compare/v1.2.0...v1.3.0) (2026-07-21)


### Features

* **brokered-connection:** wire data-source create/read to brokered_connection_id ([#21699](https://github.com/run-llama/llama-parse-go/issues/21699)) ([42370a3](https://github.com/run-llama/llama-parse-go/commit/42370a3695a6e6b8b4386e35771a476481def2eb))
* **gdrive:** reuse-first connection picker in the data-source connect modal ([#21725](https://github.com/run-llama/llama-parse-go/issues/21725)) ([8f68bd6](https://github.com/run-llama/llama-parse-go/commit/8f68bd602f5bb34af1eb49f37ed7923323312e82))
* **llamaparse:** agentic 2026-07-15 — Markdown-pipe table body for Gemini 3.1 Flash-Lite (EU primary) ([#22208](https://github.com/run-llama/llama-parse-go/issues/22208)) ([93c71a1](https://github.com/run-llama/llama-parse-go/commit/93c71a18cf6dd97485e3627a2c8aa2e099ede4b9))
* **parse:** adding forms pass to api layer (forms=`enrich` param + output types) ([#22012](https://github.com/run-llama/llama-parse-go/issues/22012)) ([a971e2a](https://github.com/run-llama/llama-parse-go/commit/a971e2a0a70b6075440a53d383168bb53edac05e))
* **parse:** confidence_scores="verified" — per-page AI-verified confidence + document-level score ([#22083](https://github.com/run-llama/llama-parse-go/issues/22083)) ([6b81de6](https://github.com/run-llama/llama-parse-go/commit/6b81de62a6621682eb6b5ef5154b147738398b49))
* **parse:** rename confidence scoring option + billing event (confidence_score_effort / confidence_score_high) ([#22290](https://github.com/run-llama/llama-parse-go/issues/22290)) ([9aca741](https://github.com/run-llama/llama-parse-go/commit/9aca741eeab198c0dca8258f3daf6649f21ce657))
* **sdk:** drop the `prod` project suffix from Go and Java namespaces ([868a030](https://github.com/run-llama/llama-parse-go/commit/868a030bf2f89ea49fab2055675cc67ddc699c2a))


### Bug Fixes

* **parse:** declare the recursive form node schemas as models ([d450216](https://github.com/run-llama/llama-parse-go/commit/d450216a1a7bda087f216ffe35ca3c8b1ed7fd8e))

## [1.2.0](https://github.com/run-llama/llama-parse-go/compare/v1.1.0...v1.2.0) (2026-07-09)


### Features

* **agentic-plus:** dated version 2026-07-08 — graduate decomposed-gemini (flash-lite), fallback to 2026-06-18 ([#21738](https://github.com/run-llama/llama-parse-go/issues/21738)) ([ac24c9d](https://github.com/run-llama/llama-parse-go/commit/ac24c9db7a3676cc5db23cf5601c4a6ebada4340))
* update fast tier latest version to use liteparse + markdown ([#21669](https://github.com/run-llama/llama-parse-go/issues/21669)) ([2df5fa2](https://github.com/run-llama/llama-parse-go/commit/2df5fa2abf6342214f5729cf28fc344371c17e6c))

## [1.1.0](https://github.com/run-llama/llama-parse-go/compare/v1.0.0...v1.1.0) (2026-06-30)


### Features

* **index:** add output_directory_id to IndexResponse ([#21149](https://github.com/run-llama/llama-parse-go/issues/21149)) ([bed12f3](https://github.com/run-llama/llama-parse-go/commit/bed12f31a6f7c8bbbadb38a6e43f959435fdd459))


### Bug Fixes

* **ci:** keep staging workflows on back-sync so GITHUB_TOKEN can push ([210783b](https://github.com/run-llama/llama-parse-go/commit/210783b00e32285d6e05fbddf72cf36d43dc6f73))
* **ci:** merge back-sync PR immediately instead of --auto ([883b7a7](https://github.com/run-llama/llama-parse-go/commit/883b7a7650ada825601728dc1aee96a172ae41ca))
* **ci:** squash-MERGE production into staging, don't reset to it ([0729cc5](https://github.com/run-llama/llama-parse-go/commit/0729cc5744ec0c22428133a0bc4e8c645cef9011))

## [1.0.0](https://github.com/run-llama/llama-parse-go/compare/v0.0.1...v1.0.0) (2026-06-25)


### Features

* **api:** Adding in Resources For Retrieval and Chat ([7f89b7b](https://github.com/run-llama/llama-parse-go/commit/7f89b7bff65d5d638899bad3154ca7ef9f83238f))
* **api:** api update ([49e4a47](https://github.com/run-llama/llama-parse-go/commit/49e4a475c172faac71c842aad7272de3722f6bd4))
* **api:** api update ([8f283a9](https://github.com/run-llama/llama-parse-go/commit/8f283a9d4312183a9d570c458c7e6127d4d0c912))
* **api:** api update ([22b0671](https://github.com/run-llama/llama-parse-go/commit/22b067153aa1c6d05a5727bab473cd70f7cefcf1))
* **api:** api update ([2618a02](https://github.com/run-llama/llama-parse-go/commit/2618a0265ff3ad850e0efdc225bcc8b904ac8dd4))
* **api:** api update ([afd84a2](https://github.com/run-llama/llama-parse-go/commit/afd84a2188fe6550670ae33a16041113be8fd1d1))
* **api:** api update ([331d0f7](https://github.com/run-llama/llama-parse-go/commit/331d0f735bd638667e83aafe974ee5fb7f19ef34))
* **api:** api update ([fc00ac5](https://github.com/run-llama/llama-parse-go/commit/fc00ac594c74d6a5211ffdc56a5e0ebab05df956))
* **api:** api update ([2e55130](https://github.com/run-llama/llama-parse-go/commit/2e5513094703c56add36ef9f8b537d613ff376a6))
* **api:** api update ([c5361d3](https://github.com/run-llama/llama-parse-go/commit/c5361d36176bb3ed07e63bc8099e453ae477db4b))
* **api:** api update ([46e62e1](https://github.com/run-llama/llama-parse-go/commit/46e62e12cd977f3543a8a3221a7e41a0f65e0f10))
* **api:** api update ([050c292](https://github.com/run-llama/llama-parse-go/commit/050c2924d027ccd71ccd7308c21e7044b5fa0779))
* **api:** api update ([295cc6c](https://github.com/run-llama/llama-parse-go/commit/295cc6cbc55967d7c1c4fe8b78c92731291748d9))
* **api:** api update ([ede0040](https://github.com/run-llama/llama-parse-go/commit/ede0040818e6be800100fa46e2dfce4bf550161e))
* **api:** api update ([2d66be8](https://github.com/run-llama/llama-parse-go/commit/2d66be8e3eb23f37736fa57c80e8455be6436871))
* **api:** batches support ([91aa9a8](https://github.com/run-llama/llama-parse-go/commit/91aa9a8926f5a51d1580750d93311c7e84c47d28))
* **api:** swap grep and file search to grep/find + pagination ([8510bb4](https://github.com/run-llama/llama-parse-go/commit/8510bb4c626b3cb44b2a578bc0a7dedf46e41bde))
* **api:** Updating indexes endpoints within the sub resources. ([fe82f63](https://github.com/run-llama/llama-parse-go/commit/fe82f636b3469da75afd3b1ee07950219874cb9e))
* **client:** optimize json encoder for internal types ([29e8d5d](https://github.com/run-llama/llama-parse-go/commit/29e8d5d3e4f295c58b217135a92e152cc8551911))
* initial stlc build ([ebff381](https://github.com/run-llama/llama-parse-go/commit/ebff38196870419a95db273c7d7ab87cbb6b3c92))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([381c6d1](https://github.com/run-llama/llama-parse-go/commit/381c6d16cc13c996d6e6ba8a0af018a0036f255a))


### Chores

* avoid embedding reflect.Type for dead code elimination ([60aba76](https://github.com/run-llama/llama-parse-go/commit/60aba7629f3605a7a52b2a199803d93e17874a3a))
* configure new SDK language ([2a7ed44](https://github.com/run-llama/llama-parse-go/commit/2a7ed4451625afd581ef7133f8d5cac3463fc9e9))
* redact api-key headers in debug logs ([9d873b6](https://github.com/run-llama/llama-parse-go/commit/9d873b677ca5923e42423d90d54948c7c3e136aa))
* release 1.0.0 ([b3919ad](https://github.com/run-llama/llama-parse-go/commit/b3919ad94c05c9e6c34041f75bc44196b7bfb7a0))
* **spec:** sync OpenAPI spec from platform ([b08edad](https://github.com/run-llama/llama-parse-go/commit/b08edad107465514f838877789e1e24bbc8b1b8a))
* update SDK settings ([1654a7f](https://github.com/run-llama/llama-parse-go/commit/1654a7f126369c8500a8c8ce2dd2e5c6bba3927f))
