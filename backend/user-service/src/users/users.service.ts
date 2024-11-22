import { Body, Injectable, NotFoundException } from '@nestjs/common';
import { User } from './user.model';
import { BanUserDto } from './dto/banUser.dto';
import { InjectModel } from '@nestjs/sequelize';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
import { AddAndDeleteFavBooks } from './dto/addAndDeleteBook.dto';

@Injectable()
export class UsersService {
  constructor(@InjectModel(User) private readonly userRepo: typeof User) {}
  async banUser(@Body() dto: BanUserDto) {
    const banUser = await this.userRepo.update(
      {
        isBanned: true,
      },
      {
        where: {
          id: dto.id,
        },
      },
    );
    if (banUser[dto.id] == 0) {
      return new NotFoundException(`User with ${dto.id} not found`).getStatus;
    }
    return { msg: `User with ID ${dto.id} has been banned` };
  }
  async unBanUser(@Body() dto: UnBanUserDto) {
    const unBanUser = await this.userRepo.update(
      {
        isBanned: false,
      },
      {
        where: {
          id: dto.id,
        },
      },
    );
    if (unBanUser[dto.id] == 0) {
      return new NotFoundException(`User with ${dto.id} not found`).getStatus;
    }
    return { msg: `User with ID ${dto.id} has been unbanned` };
  }
  async getUser(
    @Body()
    payload: any,
  ) {
    try {
      return payload
    } catch (err) {
      console.log(`ошибка: ${err}`);
    }
  }
  async changeEmail(@Body() dto: ChangeEmailDto) {
    const email = await this.userRepo.update(
      {
        email: dto.newEmail,
      },
      {
        where: {
          id: dto.userId,
        },
      },
    );
    if (email[dto.userId] == 0) {
      return new NotFoundException(`User with ${dto.userId} not found`)
        .getStatus;
    }
    return email;
  }

  async changeName(@Body() dto: ChangeNameDto) {
    const username = await this.userRepo.update(
      {
        username: dto.newUsername,
      },
      {
        where: {
          id: dto.userId,
        },
      },
    );
    if (username[dto.userId] == 0) {
      return new NotFoundException(`User with ${dto.userId} not found`)
        .getStatus;
    }
    return username;
  }
  async changeAvatarUrl(@Body() dto: changeAvatarUrlDto) {
    const AvatarUrl = await this.userRepo.update(
      {
        avatarFileUrl: dto.newAvatarUrl,
      },
      {
        where: {
          id: dto.id,
        },
      },
    );
    if (AvatarUrl[dto.id] == 0) {
      return new NotFoundException(`User with ${dto.id} not found`).getStatus;
    }
    return AvatarUrl;
  }
  async changeDescription(@Body() dto: changeDescriptionDto) {
    const description = await this.userRepo.update(
      {
        description: dto.newDescription,
      },
      {
        where: {
          id: dto.userId,
        },
      },
    );
    if (description[dto.userId] == 0) {
      return new NotFoundException(`User with ${dto.userId} not found`)
        .getStatus;
    }
    return description;
  }
  async changeFavouriteBooks(@Body() dto: AddAndDeleteFavBooks) {
    const addBooks = await this.userRepo.update(
      {
        ListBook: dto.ListBooks,
      },
      {
        where: {
          id: dto.id,
        },
      },
    );
    return addBooks;
  }
}
