import { UsersService } from './users.service';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
import { Request } from 'express';
import { AddAndDeleteFavBooks } from './dto/addAndDeleteBook.dto';
export declare class UsersController {
    private readonly usersService;
    constructor(usersService: UsersService);
    banUser(dto: BanUserDto): Promise<(() => number) | {
        msg: string;
    }>;
    getUser(req: Request): Promise<any>;
    changeFavouriteBooks(dto: AddAndDeleteFavBooks): Promise<any>;
    changeEmail(dto: ChangeEmailDto): Promise<(() => number) | [affectedCount: number]>;
    unBanUser(dto: UnBanUserDto): Promise<(() => number) | {
        msg: string;
    }>;
    changeName(dto: ChangeNameDto): Promise<any>;
    changeDescription(dto: changeDescriptionDto): Promise<any>;
    changeAvatarUrl(dto: changeAvatarUrlDto): Promise<any>;
}
