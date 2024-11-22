import { User } from './user.model';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
import { AddAndDeleteFavBooks } from './dto/addAndDeleteBook.dto';
export declare class UsersService {
    private readonly userRepo;
    constructor(userRepo: typeof User);
    banUser(dto: BanUserDto): Promise<(() => number) | {
        msg: string;
    }>;
    unBanUser(dto: UnBanUserDto): Promise<(() => number) | {
        msg: string;
    }>;
    getUser(payload: any): Promise<any>;
    changeEmail(dto: ChangeEmailDto): Promise<[affectedCount: number] | (() => number)>;
    changeName(dto: ChangeNameDto): Promise<[affectedCount: number] | (() => number)>;
    changeAvatarUrl(dto: changeAvatarUrlDto): Promise<[affectedCount: number] | (() => number)>;
    changeDescription(dto: changeDescriptionDto): Promise<[affectedCount: number] | (() => number)>;
    changeFavouriteBooks(dto: AddAndDeleteFavBooks): Promise<[affectedCount: number]>;
}
