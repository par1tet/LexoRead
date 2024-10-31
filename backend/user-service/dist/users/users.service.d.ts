import { User } from './user.model';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { GetUserDto } from './dto/getUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
import { JwtPayload } from 'jwt-decode';
export declare class UsersService {
    private userRepo;
    constructor(userRepo: typeof User);
    banUser(dto: BanUserDto): Promise<() => number>;
    unBanUser(dto: UnBanUserDto): Promise<() => number>;
    getUser(dto: GetUserDto): Promise<JwtPayload>;
    changeEmail(dto: ChangeEmailDto): Promise<[affectedCount: number] | (() => number)>;
    changeName(dto: ChangeNameDto): Promise<[affectedCount: number] | (() => number)>;
    changeAvatarUrl(dto: changeAvatarUrlDto): Promise<[affectedCount: number] | (() => number)>;
    changeDescription(dto: changeDescriptionDto): Promise<[affectedCount: number] | (() => number)>;
}
