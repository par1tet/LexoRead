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
exports.FavBooks = void 0;
const sequelize_1 = require("sequelize");
const sequelize_typescript_1 = require("sequelize-typescript");
const user_model_1 = require("./user.model");
const swagger_1 = require("@nestjs/swagger");
let FavBooks = class FavBooks extends sequelize_typescript_1.Model {
};
exports.FavBooks = FavBooks;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 1, description: 'Уникальный индификатор' }),
    (0, sequelize_typescript_1.Column)({ type: sequelize_typescript_1.DataType.INTEGER, autoIncrement: true, primaryKey: true }),
    __metadata("design:type", Number)
], FavBooks.prototype, "id", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({
        example: [1, 2, 3, 4, 5],
        description: 'получение книг по Id',
    }),
    (0, sequelize_typescript_1.Column)({ type: sequelize_typescript_1.DataType.ARRAY(sequelize_1.INTEGER), allowNull: true }),
    __metadata("design:type", Number)
], FavBooks.prototype, "List", void 0);
__decorate([
    (0, sequelize_typescript_1.BelongsTo)(() => user_model_1.User),
    __metadata("design:type", user_model_1.User)
], FavBooks.prototype, "user", void 0);
__decorate([
    (0, sequelize_typescript_1.ForeignKey)(() => user_model_1.User),
    __metadata("design:type", Number)
], FavBooks.prototype, "user_id", void 0);
exports.FavBooks = FavBooks = __decorate([
    (0, sequelize_typescript_1.Table)({ tableName: 'booksFavorite' })
], FavBooks);
//# sourceMappingURL=favBooks.model.js.map