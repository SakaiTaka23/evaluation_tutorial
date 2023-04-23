# nest

Nest.jsでのOpen Telemetryの使用例
参考 https://www.tomray.dev/nestjs-open-telemetry

まずは設定ファイルを作成する`tracing.ts`
spanProcessor → トレースの方法
instrumentations → 自動トレースのライブラリを入れる
[自動で入れることのできる物のリスト](https://opentelemetry.io/ecosystem/registry/?language=js&component=instrumentation)
設定ファイルができたら起動時の**一番最初**に関数を配置する

なぜか`@opentelemetry/api`がインストールされないので手動で追加

複数のサービスからトレースデータを集めている時でも区別をしやすくするためにリソースに対して名前を付ける方が良い

本番環境などでトレースの数が多い場合バッチ処理を行うことによって負荷を回避する
`SimpleSpanProcessor` → スパンを直接exportorに渡す
`BatchSpanProcessor` → 終了したスパンをバッチ的に渡す パフォーマンスや最適化を行う場合はこちらを使用する
