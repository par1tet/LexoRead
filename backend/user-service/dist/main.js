"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const core_1 = require("@nestjs/core");
const app_module_1 = require("./app.module");
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const helmet_1 = require("helmet");
async function bootstrap() {
    const app = await core_1.NestFactory.create(app_module_1.AppModule);
    app.setGlobalPrefix('LexoRead');
    app.enableCors();
    app.use((0, helmet_1.default)({
        crossOriginEmbedderPolicy: false,
        contentSecurityPolicy: {
            directives: {
                imgSrc: [
                    `'self'`,
                    'data:',
                    'apollo-server-landing-page.cdn.apollographql.com',
                ],
                scriptSrc: [`'self'`, `https: 'unsafe-inline'`],
                manifestSrc: [
                    `'self'`,
                    'apollo-server-landing-page.cdn.apollographql.com',
                ],
                frameSrc: [`'self'`, 'sandbox.embed.apollographql.com'],
            },
        },
    }));
    const config = new swagger_1.DocumentBuilder()
        .setTitle('Users')
        .addBearerAuth()
        .addSecurityRequirements('Bearer')
        .setDescription('The Users Api Documentation')
        .setVersion('1.0')
        .build();
    const document = swagger_1.SwaggerModule.createDocument(app, config);
    swagger_1.SwaggerModule.setup('LexoRead', app, document);
    const PORT = process.env.PORT;
    console.log(PORT);
    if (!app) {
        return new common_1.NotFoundException('NOT FOUND').getStatus();
    }
    await app.listen(PORT, () => {
        console.log(`Server is starting ${PORT}`);
    });
}
bootstrap();
//# sourceMappingURL=main.js.map