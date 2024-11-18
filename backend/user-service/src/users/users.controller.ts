import {
    Body,
    Controller,
    Delete,
    Post,
    Put,
    Res,
} from '@nestjs/common';
import {UsersService} from './users.service';
import {BanUserDto} from './dto/banUser.dto';
import {UnBanUserDto} from './dto/unBanUser.dto';
import {GetUserDto} from './dto/getUser.dto';
import {ChangeEmailDto} from './dto/changeEmail.dto';
import {ChangeNameDto} from './dto/changeName.dto';
import {changeAvatarUrlDto} from './dto/changeAvatarUrl.dto';
import {changeDescriptionDto} from './dto/changeDescription.dto';
import {
    ApiCreatedResponse,
    ApiOperation,
    ApiResponse,
    ApiTags,
} from '@nestjs/swagger';
import {User} from './user.model';
import {Response} from 'express';

@Controller('users')
@ApiTags('users')
export class UsersController {
    constructor(private readonly usersService: UsersService) {
    }

    @Put('banUser')
    @ApiOperation({summary: 'забанить пользователя'})
    async banUser(@Body() dto: BanUserDto) {
        return await this.usersService.banUser(dto);
    }

    @Post('getUser')
    @ApiOperation({summary: 'получить пользователя'})
    @ApiResponse({status: 200, description: 'get all', type: User})
    async getUser(@Body() dto: GetUserDto, @Res() response: Response) {
        const user = await this.usersService.getUser(dto);
        return response.status(201).json(user);
    }

    @Delete('deleteBook')
    @ApiOperation({summary: 'удалить книгу пользователя'})
    async deleteBook() {
    }

    @Put('changeFavBooks')
    @ApiOperation({summary: 'изменить любимые книги пользователя'})
    async changeFavouriteBooks() {
    }

    @Put('changeEmail')
    @ApiOperation({summary: 'изменить почту пользователя'})
    async changeEmail(@Body() dto: ChangeEmailDto, @Res() response: Response) {
        const email = await this.usersService.changeEmail(dto);
        return response.status(201).json(email)
    }

    @Put('unBanUser')
    @ApiOperation({summary: 'разбанить пользователя'})
    async unBanUser(@Body() dto: UnBanUserDto) {
        return await this.usersService.unBanUser(dto);
    }

    @Put('changeName')
    @ApiOperation({summary: 'изменить имя пользователя'})
    async changeName(@Body() dto: ChangeNameDto) {
        return await this.usersService.changeName(dto);
    }

    @Put('changeDescription')
    @ApiOperation({summary: 'изменить описание пользователя'})
    async changeDescription(@Body() dto: changeDescriptionDto) {
        return await this.usersService.changeDescription(dto);
    }

    @ApiCreatedResponse({
        description: 'The record has been successfully created.',
        type: User,
        status: 200,
    })
    @Put('changeAvatarUrl')
    @ApiOperation({summary: 'изменить аватарку пользователя'})
    async changeAvatarUrl(@Body() dto: changeAvatarUrlDto) {
        return await this.usersService.changeAvatarUrl(dto);
    }
}
