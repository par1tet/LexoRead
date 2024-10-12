import { UsersService } from './users.service';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { GetUserDto } from './dto/getUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
export declare class UsersController {
    private readonly usersService;
    constructor(usersService: UsersService);
    banUser(dto: BanUserDto): Promise<() => number>;
    getUser(dto: GetUserDto): Promise<import("./user.model").User>;
    deleteBook(): Promise<void>;
    changeFavouriteBooks(): Promise<void>;
    changeEmail(dto: ChangeEmailDto): Promise<[affectedCount: number] | (() => number)>;
    unBanUser(dto: UnBanUserDto): Promise<() => number>;
    changeName(dto: ChangeNameDto): Promise<[affectedCount: number] | (() => number)>;
    changeDescription(dto: changeDescriptionDto): Promise<[affectedCount: number] | (() => number)>;
    changeAvatarUrl(dto: changeAvatarUrlDto): Promise<[affectedCount: number] | (() => number)>;
}
