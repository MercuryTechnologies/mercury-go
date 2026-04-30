# Changelog

## 0.6.1 (2026-04-30)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/MercuryTechnologies/mercury-go/compare/v0.6.0...v0.6.1)

### Chores

* avoid embedding reflect.Type for dead code elimination ([45697de](https://github.com/MercuryTechnologies/mercury-go/commit/45697de9c6a053bbffcdfe57e039ac25d012fd5f))
* update OpenAPI spec from 915a0ba3e3f52bbd165420618396b2fd674694d6 ([f3012c2](https://github.com/MercuryTechnologies/mercury-go/commit/f3012c2bdc9b56ca76d9d3a3556e32d5a609fdee))

## 0.6.0 (2026-04-27)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/MercuryTechnologies/mercury-go/compare/v0.5.0...v0.6.0)

### Features

* support setting headers via env ([5b49280](https://github.com/MercuryTechnologies/mercury-go/commit/5b49280b1a83155d583e47834b4fb3f67a304c31))

## 0.5.0 (2026-04-27)

Full Changelog: [v0.4.5...v0.5.0](https://github.com/MercuryTechnologies/mercury-go/compare/v0.4.5...v0.5.0)

### Features

* **go:** add default http client with timeout ([dfff2c6](https://github.com/MercuryTechnologies/mercury-go/commit/dfff2c642ac4e1db63e7a8058888d193f8a40aac))

## 0.4.5 (2026-04-23)

Full Changelog: [v0.4.4...v0.4.5](https://github.com/MercuryTechnologies/mercury-go/compare/v0.4.4...v0.4.5)

### Chores

* update OpenAPI spec from 08d68d0d6fb7730b5df73bc1a3513ba87e2dff2f ([5bade18](https://github.com/MercuryTechnologies/mercury-go/commit/5bade18c43357b3db84e0184711366aaebe24def))

## 0.4.4 (2026-04-22)

Full Changelog: [v0.4.3...v0.4.4](https://github.com/MercuryTechnologies/mercury-go/compare/v0.4.3...v0.4.4)

### Chores

* **internal:** more robust bootstrap script ([584df4b](https://github.com/MercuryTechnologies/mercury-go/commit/584df4b032d7738a6a77e971c1f9c4116e969253))

## 0.4.3 (2026-04-13)

Full Changelog: [v0.4.2...v0.4.3](https://github.com/MercuryTechnologies/mercury-go/compare/v0.4.2...v0.4.3)

### Refactors

* rename list_transactions -&gt; transactions ([e414464](https://github.com/MercuryTechnologies/mercury-go/commit/e41446445f20863956456a3af788d44c6f7c4a7d))

## 0.4.2 (2026-04-13)

Full Changelog: [v0.4.1...v0.4.2](https://github.com/MercuryTechnologies/mercury-go/compare/v0.4.1...v0.4.2)

### Chores

* remove old list_attachments. ([c67478c](https://github.com/MercuryTechnologies/mercury-go/commit/c67478c5724effbb86ff2691c2ac00e00a603f83))


### Refactors

* make attachments a sub resource ([4629e8e](https://github.com/MercuryTechnologies/mercury-go/commit/4629e8edafd5840ebc2f7b3b908a5100e6c33ddb))

## 0.4.1 (2026-04-13)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/MercuryTechnologies/mercury-go/compare/v0.4.0...v0.4.1)

### Refactors

* consolidate money movement into single payments resource ([f6d9145](https://github.com/MercuryTechnologies/mercury-go/commit/f6d91452184a57d354629920d21f30cbaebcf5ae))
* move ar attachments as a sub resrource ([7491d94](https://github.com/MercuryTechnologies/mercury-go/commit/7491d94868b5fed08a38701e29b4176658bd9c5f))
* rename upload_attachment -&gt; attach ([2f5cbf2](https://github.com/MercuryTechnologies/mercury-go/commit/2f5cbf2e4b839157d73e9747b81d2215d82af28c))
* unify statements into single resource with account/treasury subresources ([37db929](https://github.com/MercuryTechnologies/mercury-go/commit/37db929ec3c57a94ff503db0cd61e4087b1ce62b))
* unify transactions listing into single resource ([60f504c](https://github.com/MercuryTechnologies/mercury-go/commit/60f504c3d347ad593fdb5491efe395d44ac39561))

## 0.4.0 (2026-04-09)

Full Changelog: [v0.3.3...v0.4.0](https://github.com/MercuryTechnologies/mercury-go/compare/v0.3.3...v0.4.0)

### ⚠ BREAKING CHANGES

* remove account.get_transaction

### Features

* remove account.get_transaction ([731aeed](https://github.com/MercuryTechnologies/mercury-go/commit/731aeedce97fb40adfb736872f70b6f95c3648b7))


### Styles

* use `download` as method ([b024e06](https://github.com/MercuryTechnologies/mercury-go/commit/b024e062478cae5a2fec500d820b17ba6a326c0a))

## 0.3.3 (2026-04-09)

Full Changelog: [v0.3.2...v0.3.3](https://github.com/MercuryTechnologies/mercury-go/compare/v0.3.2...v0.3.3)

### Bug Fixes

* get events typo ([a5d7f04](https://github.com/MercuryTechnologies/mercury-go/commit/a5d7f04d6377f300d5d2cad4511ba0947c205d49))


### Chores

* fix attachments typo ([ed13fb4](https://github.com/MercuryTechnologies/mercury-go/commit/ed13fb4e639f552b8a9da0800b3b686d5bde05cf))


### Styles

* rename "retrieve" to "get" ([828c52c](https://github.com/MercuryTechnologies/mercury-go/commit/828c52c39b4efa5ce98a29bd604a5772eae898cf))
* rename `organization` to `org` ([849dc13](https://github.com/MercuryTechnologies/mercury-go/commit/849dc1374beb10e96bde4d9f84e8bc0855a5a9d6))


### Refactors

* create cards resource ([0a93dac](https://github.com/MercuryTechnologies/mercury-go/commit/0a93dac162322f7c47c031ee611aca4d9971a557))
* make `customers` top level resource ([5fe90f0](https://github.com/MercuryTechnologies/mercury-go/commit/5fe90f01286088d88479f7830e7340db57a7e9f2))
* make `invoices` top level resource ([7a2ad26](https://github.com/MercuryTechnologies/mercury-go/commit/7a2ad26a78edaa0a25412743c476002c6044a3df))

## 0.3.2 (2026-04-08)

Full Changelog: [v0.3.1...v0.3.2](https://github.com/MercuryTechnologies/mercury-go/compare/v0.3.1...v0.3.2)

### Bug Fixes

* remove basicAuth ([392d103](https://github.com/MercuryTechnologies/mercury-go/commit/392d103349474bbf77f7592f490c371e55768f84))

## 0.3.1 (2026-04-08)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/MercuryTechnologies/mercury-go/compare/v0.3.0...v0.3.1)

### Chores

* update OpenAPI spec from 28e7cd738b9c28911bdd51b6d64e7036d42b8f8e ([a645cfa](https://github.com/MercuryTechnologies/mercury-go/commit/a645cfadedaee32f3cb06b97d3db3f58f069582c))

## 0.3.0 (2026-04-08)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/MercuryTechnologies/mercury-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** fix sendmoneyrequest id pagination ([414fa28](https://github.com/MercuryTechnologies/mercury-go/commit/414fa289589ebde3cf62df258a4fc8b8966747b6))
* **api:** remove mcp generation ([6967067](https://github.com/MercuryTechnologies/mercury-go/commit/6967067b18c650c093b11b6ae675be8a14f88709))


### Bug Fixes

* force v0.2.0 to unblock build ([d5e6300](https://github.com/MercuryTechnologies/mercury-go/commit/d5e6300384076708b31a2f0ec7df96520cffde16))


### Chores

* sync repo ([29345c3](https://github.com/MercuryTechnologies/mercury-go/commit/29345c33ccb31f1d1df430e76a7a5f3561e8071f))
* update SDK settings ([d554657](https://github.com/MercuryTechnologies/mercury-go/commit/d55465749dcca49f95845bc9bc89e5b156983ec7))

## 0.1.0 (2026-03-31)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/MercuryTechnologies/mercury-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** api update ([bfd79c0](https://github.com/MercuryTechnologies/mercury-go/commit/bfd79c0d61f0c14d0925bfbe2b12f10451c4ce6e))
* **api:** api update ([8dd50f2](https://github.com/MercuryTechnologies/mercury-go/commit/8dd50f27a3729b9b57a79c6c575c5f85b44a51a5))
* **api:** api update ([095069d](https://github.com/MercuryTechnologies/mercury-go/commit/095069df389bf720e78af60ed6a8177961b17263))
* **api:** api update ([82060c7](https://github.com/MercuryTechnologies/mercury-go/commit/82060c77e7ae6256fdf0b30883b5f2199e843c4b))
* **api:** api update ([9a48a98](https://github.com/MercuryTechnologies/mercury-go/commit/9a48a98c13cf5cccf4b813ca897b1e2de51cf64e))
* **api:** api update ([e5c951a](https://github.com/MercuryTechnologies/mercury-go/commit/e5c951a1178b48ce35b15a5ba33c64c7320d1a53))
* **api:** api update ([ac8d9e3](https://github.com/MercuryTechnologies/mercury-go/commit/ac8d9e35bd308856737cb112d915927d2b75618d))
* **api:** api update ([3030512](https://github.com/MercuryTechnologies/mercury-go/commit/303051230ec0b6e7d74e99685ca3fe1ec440d5aa))
* **api:** api update ([679b5ab](https://github.com/MercuryTechnologies/mercury-go/commit/679b5ab24b36072c70703d847deb3669a706947a))
* **api:** api update ([36671f3](https://github.com/MercuryTechnologies/mercury-go/commit/36671f3851e359a8d176c07469aa3f94e6ca8ea6))
* **api:** api update ([8ce3e77](https://github.com/MercuryTechnologies/mercury-go/commit/8ce3e7726e04664ee86babf45cf07acb296c1374))
* **api:** api update ([89f0407](https://github.com/MercuryTechnologies/mercury-go/commit/89f04078d0f63c84df1788e934ebd5dc02fcd361))
* **api:** api update ([ed6e40e](https://github.com/MercuryTechnologies/mercury-go/commit/ed6e40ee8565a24eccf21ebe9e0deacfbdac93a5))
* **api:** api update ([e4a5cb2](https://github.com/MercuryTechnologies/mercury-go/commit/e4a5cb2406054230c7e0d9fed86d1d66f5b3a79a))
* **api:** api update ([8a46f3d](https://github.com/MercuryTechnologies/mercury-go/commit/8a46f3d6484513f7e30fdf3646ec02272ff520e4))
* **api:** api update ([60bfefe](https://github.com/MercuryTechnologies/mercury-go/commit/60bfefe0f51cc4265e6fc8af9125cc80e864962e))
* **api:** api update ([79e0bd5](https://github.com/MercuryTechnologies/mercury-go/commit/79e0bd5ddc2e2920afa853aea0f12ac3639f0bdc))
* **api:** api update ([e5b926c](https://github.com/MercuryTechnologies/mercury-go/commit/e5b926c6b13350543339cfb92b4fc0e3569de1b3))
* **api:** api update ([8f99fe3](https://github.com/MercuryTechnologies/mercury-go/commit/8f99fe33ea7a2fe4db97ae198f5ef38f83d710d7))
* **api:** manual updates ([6ad332e](https://github.com/MercuryTechnologies/mercury-go/commit/6ad332e2c8f45828a3cb1f35528cd565a33e4e34))
* **api:** manual updates ([e693458](https://github.com/MercuryTechnologies/mercury-go/commit/e6934586cdea91a263d1f4753d5b4989aa3c839a))
* **api:** manual updates ([eda3817](https://github.com/MercuryTechnologies/mercury-go/commit/eda3817398ee14fb19fdc3d1546f1f35ff351654))
* **api:** manual updates ([7af09cc](https://github.com/MercuryTechnologies/mercury-go/commit/7af09ccb8d5e79d8206af041e9d679e2a96d6b7b))
* **api:** manual updates ([15d7e76](https://github.com/MercuryTechnologies/mercury-go/commit/15d7e767d47534eca158d7be7a539e6ac79083b2))
* **api:** manual updates ([fa6b3e6](https://github.com/MercuryTechnologies/mercury-go/commit/fa6b3e60efa6297cae3636337895d4b623eff794))
* **api:** manual updates ([53384bc](https://github.com/MercuryTechnologies/mercury-go/commit/53384bc93074517cf223de2f29621268eb6a74a8))
* **internal:** support comma format in multipart form encoding ([e435c50](https://github.com/MercuryTechnologies/mercury-go/commit/e435c50d99ed0dad2b9db345fe227ebc6819a72e))


### Bug Fixes

* prevent duplicate ? in query params ([1a677db](https://github.com/MercuryTechnologies/mercury-go/commit/1a677dbc82b86707651035a199d3e1f49725838a))


### Chores

* **ci:** skip lint on metadata-only changes ([be67dd9](https://github.com/MercuryTechnologies/mercury-go/commit/be67dd95953652b58012eeb6c5db8919d98a0363))
* **ci:** skip uploading artifacts on stainless-internal branches ([8b0bffd](https://github.com/MercuryTechnologies/mercury-go/commit/8b0bffdd8b2fbc8cb219bfb374ae02cabbe32da3))
* **ci:** support opting out of skipping builds on metadata-only commits ([5a450a1](https://github.com/MercuryTechnologies/mercury-go/commit/5a450a12182c0ef8f9e645a35c0471dfeb3d8b7f))
* **client:** fix multipart serialisation of Default() fields ([8e02730](https://github.com/MercuryTechnologies/mercury-go/commit/8e02730a08a8a0696a935a9a550d8c8b6c9800d7))
* configure new SDK language ([24d7918](https://github.com/MercuryTechnologies/mercury-go/commit/24d7918f82639705c93717e28ef0dae1d11cff7b))
* **internal:** codegen related update ([dfa22ad](https://github.com/MercuryTechnologies/mercury-go/commit/dfa22ad079e03dd6954288812bb632fb3ffb08e8))
* **internal:** codegen related update ([dec6125](https://github.com/MercuryTechnologies/mercury-go/commit/dec6125a189da3e7a68aca75efc32bdd952eab8f))
* **internal:** minor cleanup ([ba745bb](https://github.com/MercuryTechnologies/mercury-go/commit/ba745bb2a5c482663d0fc8392d88974ec71d8f53))
* **internal:** move custom custom `json` tags to `api` ([52fe132](https://github.com/MercuryTechnologies/mercury-go/commit/52fe132f51b08211c5666e57b7f1b9a6f1ed8757))
* **internal:** support default value struct tag ([694940d](https://github.com/MercuryTechnologies/mercury-go/commit/694940dad5bc915af074f8ac95455466cd995325))
* **internal:** tweak CI branches ([86ce11d](https://github.com/MercuryTechnologies/mercury-go/commit/86ce11d1f91b3837a8b0bad864aeb2e58e4f557b))
* **internal:** update gitignore ([2723fae](https://github.com/MercuryTechnologies/mercury-go/commit/2723faeae4d0f3ffa988298c5ea7ac1937b3107e))
* **internal:** use explicit returns ([2f4c785](https://github.com/MercuryTechnologies/mercury-go/commit/2f4c7857cf086e5485497da5d0a79267353d09de))
* **internal:** use explicit returns in more places ([d2d59b0](https://github.com/MercuryTechnologies/mercury-go/commit/d2d59b0a6fb27bfe94c6ba25fe84d739f7c04a28))
* remove unnecessary error check for url parsing ([1316ab4](https://github.com/MercuryTechnologies/mercury-go/commit/1316ab4cefb3a518ade19aa7760109d08f2151e4))
* update docs for api:"required" ([626f8f4](https://github.com/MercuryTechnologies/mercury-go/commit/626f8f4b9d821d87a35aef5af64f926327267387))
* update OpenAPI spec from 3dde921878c399c6dde2f853c2c18d9d955d4f7d ([e67b60e](https://github.com/MercuryTechnologies/mercury-go/commit/e67b60e0ece10e704c573dfef08f49e30bc673dd))
* update OpenAPI spec from 646e2560aee76feae61bf71591fb6db08efc338f ([c466bc9](https://github.com/MercuryTechnologies/mercury-go/commit/c466bc93ef0a6db9c8273d67f716daec1461e04e))
* update OpenAPI spec from 6d1dd08c2dcffcbd3826b8e808f2dbb90305ddde ([6e7af8d](https://github.com/MercuryTechnologies/mercury-go/commit/6e7af8d1136dd93a599ba01b131fd9e521ddafb6))
* update OpenAPI spec from 72108644355aea945490035c66e83cc5f2344938 ([6840761](https://github.com/MercuryTechnologies/mercury-go/commit/6840761931b83e2f374137d204add53d3e1f63ce))
* update OpenAPI spec from a45da942ea4858e86718b1a132081277dbf81343 ([a1a2011](https://github.com/MercuryTechnologies/mercury-go/commit/a1a2011c3abb4001abfa3aaff9e85ef497261d18))
* update OpenAPI spec from ef842a3c9c4d9d635ed6ab91cc18c629b01f01b7 ([79e25e9](https://github.com/MercuryTechnologies/mercury-go/commit/79e25e98a7cfc2319f11416683cae2646f1cd5d2))
* update OpenAPI spec from v2026.03.02.13 ([dd92798](https://github.com/MercuryTechnologies/mercury-go/commit/dd9279844f38890e2f503bbb5cdc424937755976))
* update OpenAPI spec from v2026.03.03.01 ([aa91cf5](https://github.com/MercuryTechnologies/mercury-go/commit/aa91cf596dc4fcdae4d9a44fc620ae151e38a511))
* update OpenAPI spec from v2026.03.03.02 ([2fb1ef3](https://github.com/MercuryTechnologies/mercury-go/commit/2fb1ef3375f094f2a33aa7c6091e629d8d7de67a))
* update placeholder string ([b587eff](https://github.com/MercuryTechnologies/mercury-go/commit/b587eff4790a4329033e7e01917636c6e48dbc6a))
* update SDK settings ([e3149a2](https://github.com/MercuryTechnologies/mercury-go/commit/e3149a21e13b928d6ff5c07c8651e9488c60b6ff))
* update SDK settings ([dc2889a](https://github.com/MercuryTechnologies/mercury-go/commit/dc2889af40541fab3974d724a6cceca0f5ed12bf))
* update SDK settings ([22111fb](https://github.com/MercuryTechnologies/mercury-go/commit/22111fbf3b606482388792c3914095314c9ef026))
* update SDK settings ([82fd7de](https://github.com/MercuryTechnologies/mercury-go/commit/82fd7de1218b43dae8323d169e93d87f4178b8b8))
