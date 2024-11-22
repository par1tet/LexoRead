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
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.UsersController = void 0;
const common_1 = require("@nestjs/common");
const users_service_1 = require("./users.service");
const banUser_dto_1 = require("./dto/banUser.dto");
const unBanUser_dto_1 = require("./dto/unBanUser.dto");
const changeEmail_dto_1 = require("./dto/changeEmail.dto");
const changeName_dto_1 = require("./dto/changeName.dto");
const changeAvatarUrl_dto_1 = require("./dto/changeAvatarUrl.dto");
const changeDescription_dto_1 = require("./dto/changeDescription.dto");
const swagger_1 = require("@nestjs/swagger");
const user_model_1 = require("./user.model");
const jwt_decode_1 = require("jwt-decode");
const addAndDeleteBook_dto_1 = require("./dto/addAndDeleteBook.dto");
let UsersController = class UsersController {
    constructor(usersService) {
        this.usersService = usersService;
    }
    async banUser(dto) {
        return await this.usersService.banUser(dto);
    }
    async getUser(req) {
        const authHeader = req.headers.authorization;
        if (authHeader) {
            const token = authHeader.split(' ')[1];
            return this.usersService.getUser((0, jwt_decode_1.jwtDecode)(token));
        }
    }
    async changeFavouriteBooks(dto) {
        return this.usersService.changeFavouriteBooks(dto);
    }
    async changeEmail(dto) {
        return await this.usersService.changeEmail(dto);
    }
    async unBanUser(dto) {
        return await this.usersService.unBanUser(dto);
    }
    async changeName(dto) {
        return await this.usersService.changeName(dto);
    }
    async changeDescription(dto) {
        return await this.usersService.changeDescription(dto);
    }
    async changeAvatarUrl(dto) {
        return await this.usersService.changeAvatarUrl(dto);
    }
};
exports.UsersController = UsersController;
__decorate([
    (0, common_1.Put)('banUser'),
    (0, swagger_1.ApiOperation)({ summary: 'забанить пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [banUser_dto_1.BanUserDto]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "banUser", null);
__decorate([
    (0, common_1.Post)('getUser'),
    (0, swagger_1.ApiOperation)({ summary: 'получить пользователя' }),
    __param(0, (0, common_1.Req)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "getUser", null);
__decorate([
    (0, common_1.Put)('changeFavBooks'),
    (0, swagger_1.ApiOperation)({ summary: 'изменить любимые книги пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [addAndDeleteBook_dto_1.AddAndDeleteFavBooks]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "changeFavouriteBooks", null);
__decorate([
    (0, swagger_1.ApiCreatedResponse)({
        description: 'The record has been successfully created.',
        type: user_model_1.User,
    }),
    (0, common_1.Put)('changeEmail'),
    (0, swagger_1.ApiOperation)({ summary: 'изменить почту пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeEmail_dto_1.ChangeEmailDto]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "changeEmail", null);
__decorate([
    (0, common_1.Put)('unBanUser'),
    (0, swagger_1.ApiOperation)({ summary: 'разбанить пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [unBanUser_dto_1.UnBanUserDto]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "unBanUser", null);
__decorate([
    (0, common_1.Put)('changeName'),
    (0, swagger_1.ApiOperation)({ summary: 'изменить имя пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeName_dto_1.ChangeNameDto]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "changeName", null);
__decorate([
    (0, common_1.Put)('changeDescription'),
    (0, swagger_1.ApiOperation)({ summary: 'изменить описание пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeDescription_dto_1.changeDescriptionDto]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "changeDescription", null);
__decorate([
    (0, swagger_1.ApiCreatedResponse)({
        description: 'The record has been successfully created.',
        type: user_model_1.User,
    }),
    (0, common_1.Put)('changeAvatarUrl'),
    (0, swagger_1.ApiOperation)({ summary: 'изменить аватарку пользователя' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeAvatarUrl_dto_1.changeAvatarUrlDto]),
    __metadata("design:returntype", Promise)
], UsersController.prototype, "changeAvatarUrl", null);
exports.UsersController = UsersController = __decorate([
    (0, swagger_1.ApiBearerAuth)(),
    (0, common_1.Controller)('users'),
    (0, swagger_1.ApiTags)('users'),
    __metadata("design:paramtypes", [users_service_1.UsersService])
], UsersController);
//# sourceMappingURL=users.controller.js.map