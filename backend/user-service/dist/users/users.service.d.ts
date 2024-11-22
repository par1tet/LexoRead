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
    changeEmail(dto: ChangeEmailDto): Promise<any>;
    changeName(dto: ChangeNameDto): Promise<any>;
    changeAvatarUrl(dto: changeAvatarUrlDto): Promise<any>;
    changeDescription(dto: changeDescriptionDto): Promise<any>;
    changeFavouriteBooks(dto: AddAndDeleteFavBooks): Promise<any>;
}
