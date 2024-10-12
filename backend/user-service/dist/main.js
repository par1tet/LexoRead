"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const core_1 = require("@nestjs/core");
const app_module_1 = require("./app.module");
const common_1 = require("@nestjs/common");
async function bootstrap() {
    const app = await core_1.NestFactory.create(app_module_1.AppModule);
    app.setGlobalPrefix('LexoRead');
    const PORT = process.env.PORT;
    if (!app) {
        return new common_1.NotFoundException('NOT FOUND').getStatus();
    }
    await app.listen(PORT, () => {
        console.log(`Server is starting ${PORT}`);
    });
}
bootstrap();
//# sourceMappingURL=main.js.map