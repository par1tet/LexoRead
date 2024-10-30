import { UsersService } from './users.service';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { GetUserDto } from './dto/getUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
import { Response } from 'express';
export declare class UsersController {
    private readonly usersService;
    constructor(usersService: UsersService);
    banUser(dto: BanUserDto): Promise<() => number>;
    getUser(dto: GetUserDto, response: Response): Promise<Response<any, Record<string, any>>>;
    deleteBook(): Promise<void>;
    changeFavouriteBooks(): Promise<void>;
    changeEmail(dto: ChangeEmailDto, response: Response): Promise<Response<any, Record<string, any>>>;
    unBanUser(dto: UnBanUserDto): Promise<() => number>;
    changeName(dto: ChangeNameDto): Promise<(() => number) | [affectedCount: number]>;
    changeDescription(dto: changeDescriptionDto): Promise<(() => number) | [affectedCount: number]>;
    changeAvatarUrl(dto: changeAvatarUrlDto): Promise<(() => number) | [affectedCount: number]>;
}
