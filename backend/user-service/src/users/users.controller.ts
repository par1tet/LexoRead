import { Body, Controller, Delete, Get, Post, Put, Req } from '@nestjs/common';
import { UsersService } from './users.service';
import { BanUserDto } from './dto/banUser.dto';
import { UnBanUserDto } from './dto/unBanUser.dto';
import { GetUserDto } from './dto/getUser.dto';
import { ChangeEmailDto } from './dto/changeEmail.dto';
import { ChangeNameDto } from './dto/changeName.dto';
import { changeAvatarUrlDto } from './dto/changeAvatarUrl.dto';
import { changeDescriptionDto } from './dto/changeDescription.dto';
import {
  ApiBearerAuth,
  ApiCreatedResponse,
  ApiExtraModels,
  ApiOperation,
  ApiParam,
  ApiTags,
} from '@nestjs/swagger';
import { User } from './user.model';
import { jwtDecode } from 'jwt-decode';
import { Request } from 'express';
@ApiBearerAuth()
@Controller('users')
@ApiTags('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Put('banUser')
  @ApiOperation({ summary: 'забанить пользователя' })
  async banUser(@Body() dto: BanUserDto) {
    return await this.usersService.banUser(dto);
  }
  @Get('getUser')
  @ApiOperation({ summary: 'получить пользователя' })
  async getUser(@Req() req: Request) {
    const authHeader = req.headers.authorization;
    if (authHeader) {
      const token = authHeader.split(' ')[1];
      const decoded = jwtDecode(token);  
      return this.usersService.getUser(decoded)
    }
  }
  @Delete('deleteBook')
  @ApiOperation({ summary: 'удалить книгу пользователя' })
  async deleteBook() {}
  @Put('changeFavBooks')
  @ApiOperation({ summary: 'изменить любимые книги пользователя' })
  async changeFavouriteBooks() {}
  @ApiCreatedResponse({
    description: 'The record has been successfully created.',
    type: User,
  })
  @Put('changeEmail')
  @ApiOperation({ summary: 'изменить почту пользователя' })
  async changeEmail(
    @Body() dto: ChangeEmailDto,
  ): Promise<(() => number) | [affectedCount: number]> {
    return await this.usersService.changeEmail(dto);
  }
  @Put('unBanUser')
  @ApiOperation({ summary: 'разбанить пользователя' })
  async unBanUser(@Body() dto: UnBanUserDto) {
    return await this.usersService.unBanUser(dto);
  }
  @Put('changeName')
  @ApiOperation({ summary: 'изменить имя пользователя' })
  async changeName(@Body() dto: ChangeNameDto) {
    return await this.usersService.changeName(dto);
  }
  @Put('changeDescription')
  @ApiOperation({ summary: 'изменить описание пользователя' })
  async changeDescription(@Body() dto: changeDescriptionDto) {
    return await this.usersService.changeDescription(dto);
  }
  @ApiCreatedResponse({
    description: 'The record has been successfully created.',
    type: User,
  })
  @Put('changeAvatarUrl')
  @ApiOperation({ summary: 'изменить аватарку пользователя' })
  async changeAvatarUrl(@Body() dto: changeAvatarUrlDto) {
    return await this.usersService.changeAvatarUrl(dto);
  }
}
