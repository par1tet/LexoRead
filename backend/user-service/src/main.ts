import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { NotFoundException } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('LexoRead');
  const config = new DocumentBuilder()
    .setTitle('Users')
    .setDescription('The Users Api Documentation')
    .setVersion('1.0')
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('LexoRead', app, document);
  const PORT = process.env.PORT;
  if (!app) {
    return new NotFoundException('NOT FOUND').getStatus();
  }
  await app.listen(PORT, () => {
    console.log(`Server is starting ${PORT}`);
  });
}
bootstrap();
