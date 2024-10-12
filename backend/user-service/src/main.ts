import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { NotFoundException } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('LexoRead')
  const PORT = process.env.PORT
  if(!app) {
    return new NotFoundException('NOT FOUND').getStatus()
  }
  await app.listen(PORT, () => {
    console.log(`Server is starting ${PORT}`);
    
  });
}
bootstrap();
