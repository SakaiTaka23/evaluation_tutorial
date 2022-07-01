import { Controller, Get } from '@nestjs/common';
import { InjectMetric } from '@willsoto/nestjs-prometheus';
import { Counter } from 'prom-client';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(
    private readonly appService: AppService,
    @InjectMetric('app_access') public counter: Counter<string>,
  ) {}

  @Get()
  getHello(): string {
    this.counter.inc();
    return this.appService.getHello();
  }
}
