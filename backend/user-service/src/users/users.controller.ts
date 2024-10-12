import { Body, Controller, Delete, Post, Put } from '@nestjs/common';
import { UsersService } from './users.service';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { GetUserDto } from './dto/getUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';

@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}
  @Put('banUser')
  async banUser(@Body() dto: BanUserDto) {
    return await this.usersService.banUser(dto)
  }
  @Post('getUser')
  async getUser(@Body() dto: GetUserDto) {
    return await this.usersService.getUser(dto)
  }
  @Delete('deleteBook')
  async deleteBook() {}
  @Put('changeFavBooks')
  async changeFavouriteBooks() {}
  @Put('changeEmail')
  async changeEmail(@Body() dto: ChangeEmailDto) {
    return await this.usersService.changeEmail(dto)
  }
  @Put('unBanUser')
  async unBanUser(@Body() dto: UnBanUserDto) {
    return await this.usersService.unBanUser(dto)
  }
  @Put('changeName')
  async changeName(@Body() dto: ChangeNameDto) {
    return await this.usersService.changeName(dto)
  }
  @Put('changeDescription')
  async changeDescription(@Body() dto: changeDescriptionDto) {
    return await this.usersService.changeDescription(dto)
  }
  @Put('changeAvatarUrl')
  async changeAvatarUrl(@Body() dto: changeAvatarUrlDto) {
    return await this.usersService.changeAvatarUrl(dto)
  }
}
