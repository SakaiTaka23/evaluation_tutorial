import { Module } from '@nestjs/common';
import {
  makeCounterProvider,
  PrometheusModule,
} from '@willsoto/nestjs-prometheus';
import { AppController } from './app.controller';
import { AppService } from './app.service';

@Module({
  imports: [PrometheusModule.register()],
  controllers: [AppController],
  providers: [
    AppService,
    makeCounterProvider({
      name: 'app_access',
      help: 'access time of app controller',
    }),
  ],
})
export class AppModule {}
