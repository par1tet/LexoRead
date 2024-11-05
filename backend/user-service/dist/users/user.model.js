"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.User = void 0;
const swagger_1 = require("@nestjs/swagger");
const sequelize_1 = require("sequelize");
const sequelize_typescript_1 = require("sequelize-typescript");
const sequelize_typescript_2 = require("sequelize-typescript");
let User = class User extends sequelize_typescript_1.Model {
};
exports.User = User;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 1, description: 'Уникальный индификатор' }),
    (0, sequelize_typescript_2.Column)({
        type: sequelize_typescript_2.DataType.INTEGER,
        primaryKey: true,
        autoIncrement: true,
        unique: true,
    }),
    __metadata("design:type", Number)
], User.prototype, "id", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: true, description: 'забанен' }),
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.BOOLEAN, defaultValue: false, allowNull: true }),
    __metadata("design:type", Boolean)
], User.prototype, "isBanned", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Artem', description: 'username' }),
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.STRING, allowNull: true }),
    __metadata("design:type", String)
], User.prototype, "username", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'smallgigached@gmail.com', description: 'email' }),
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.STRING, unique: true }),
    __metadata("design:type", String)
], User.prototype, "email", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({
        example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9',
        description: 'HashPassword',
    }),
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.STRING, allowNull: true }),
    __metadata("design:type", String)
], User.prototype, "hashPassword", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({
        example: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQy1QKMwvvFRtE7kNDp0oyv8e5k39uBdPxbbg&s',
        description: 'avatarUrl',
    }),
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.STRING(500), allowNull: true }),
    __metadata("design:type", String)
], User.prototype, "avatarFileUrl", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Hello World', description: 'description' }),
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.STRING(350) }),
    __metadata("design:type", String)
], User.prototype, "description", void 0);
__decorate([
    (0, sequelize_typescript_2.Column)({ type: sequelize_typescript_2.DataType.ARRAY(sequelize_1.INTEGER) }),
    __metadata("design:type", Array)
], User.prototype, "ListBook", void 0);
exports.User = User = __decorate([
    (0, sequelize_typescript_2.Table)({ tableName: 'users', createdAt: false, updatedAt: false })
], User);
//# sourceMappingURL=user.model.js.map