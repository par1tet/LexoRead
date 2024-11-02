"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppModule = void 0;
const common_1 = require("@nestjs/common");
const users_module_1 = require("./users/users.module");
const config_1 = require("@nestjs/config");
const sequelize_1 = require("@nestjs/sequelize");
const user_model_1 = require("./users/user.model");
const userMiddleWare_1 = require("./users/userMiddleWare");
const users_controller_1 = require("./users/users.controller");
let AppModule = class AppModule {
    configure(consumer) {
        consumer.apply(userMiddleWare_1.UserMiddleware).exclude({
            path: 'getUser',
            method: common_1.RequestMethod.POST,
        }).forRoutes(users_controller_1.UsersController);
    }
};
exports.AppModule = AppModule;
exports.AppModule = AppModule = __decorate([
    (0, common_1.Module)({
        imports: [
            config_1.ConfigModule.forRoot({
                envFilePath: '../.env',
            }),
            sequelize_1.SequelizeModule.forRoot({
                dialect: 'postgres',
                host: process.env.PG_HOST,
                port: 5432,
                username: process.env.PG_USERNAME,
                password: process.env.PG_PASSWORD,
                database: process.env.PG_DB,
                models: [user_model_1.User],
                synchronize: true,
                autoLoadModels: true,
            }),
            users_module_1.UsersModule,
        ],
        controllers: [],
        providers: [],
    })
], AppModule);
//# sourceMappingURL=app.module.js.map