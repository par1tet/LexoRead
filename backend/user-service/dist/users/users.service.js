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
exports.UsersService = void 0;
const common_1 = require("@nestjs/common");
const user_model_1 = require("./user.model");
const banUser_dto_1 = require("./dto/banUser.dto");
const sequelize_1 = require("@nestjs/sequelize");
const unBanUser_dto_1 = require("./dto/unBanUser.dto");
const getUser_dto_1 = require("./dto/getUser.dto");
const changeEmail_dto_1 = require("./dto/changeEmail.dto");
const changeName_dto_1 = require("./dto/changeName.dto");
const changeAvatarUrl_dto_1 = require("./dto/changeAvatarUrl.dto");
const changeDescription_dto_1 = require("./dto/changeDescription.dto");
let UsersService = class UsersService {
    constructor(userRepo) {
        this.userRepo = userRepo;
    }
    async banUser(dto) {
        const banUser = await this.userRepo.update({
            isBanned: true,
        }, {
            where: {
                id: dto.id,
            },
        });
        if (banUser[dto.id] == 0) {
            return new common_1.NotFoundException(`User with ${dto.id} not found`).getStatus;
        }
    }
    async unBanUser(dto) {
        const unBanUser = await this.userRepo.update({
            isBanned: false,
        }, {
            where: {
                id: dto.id,
            },
        });
        if (unBanUser[dto.id] == 0) {
            return new common_1.NotFoundException(`User with ${dto.id} not found`).getStatus;
        }
    }
    async getUser(dto) {
        const user = await this.userRepo.findOne({
            where: {
                jwtToken: dto.jwtToken,
            },
        });
        console.log(user);
        return user;
    }
    async changeEmail(dto) {
        const email = await this.userRepo.update({
            email: dto.newEmail,
        }, {
            where: {
                id: dto.userId,
            },
        });
        if (email[dto.userId] == 0) {
            return new common_1.NotFoundException(`User with ${dto.userId} not found`)
                .getStatus;
        }
        return email;
    }
    async changeName(dto) {
        const username = await this.userRepo.update({
            username: dto.newUsername,
        }, {
            where: {
                id: dto.userId,
            },
        });
        if (username[dto.userId] == 0) {
            return new common_1.NotFoundException(`User with ${dto.userId} not found`)
                .getStatus;
        }
        return username;
    }
    async changeAvatarUrl(dto) {
        const AvatarUrl = await this.userRepo.update({
            avatarFileUrl: dto.newAvatarUrl,
        }, {
            where: {
                id: dto.id,
            },
        });
        if (AvatarUrl[dto.id] == 0) {
            return new common_1.NotFoundException(`User with ${dto.id} not found`).getStatus;
        }
        return AvatarUrl;
    }
    async changeDescription(dto) {
        const description = await this.userRepo.update({
            description: dto.newDescription,
        }, {
            where: {
                id: dto.userId,
            },
        });
        if (description[dto.userId] == 0) {
            return new common_1.NotFoundException(`User with ${dto.userId} not found`)
                .getStatus;
        }
        return description;
    }
};
exports.UsersService = UsersService;
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [banUser_dto_1.BanUserDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "banUser", null);
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [unBanUser_dto_1.UnBanUserDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "unBanUser", null);
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [getUser_dto_1.GetUserDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "getUser", null);
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeEmail_dto_1.ChangeEmailDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "changeEmail", null);
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeName_dto_1.ChangeNameDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "changeName", null);
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeAvatarUrl_dto_1.changeAvatarUrlDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "changeAvatarUrl", null);
__decorate([
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [changeDescription_dto_1.changeDescriptionDto]),
    __metadata("design:returntype", Promise)
], UsersService.prototype, "changeDescription", null);
exports.UsersService = UsersService = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, sequelize_1.InjectModel)(user_model_1.User)),
    __metadata("design:paramtypes", [Object])
], UsersService);
//# sourceMappingURL=users.service.js.map