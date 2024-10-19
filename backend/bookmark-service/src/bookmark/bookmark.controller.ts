import { Body, Controller, Get, Post } from '@nestjs/common'
import { BookmarkService } from './bookmark.service'
import { CreateBookmarkDto } from './dto/create-bookmark.dto'

@Controller('bookmark')
export class BookmarkController {
	constructor(private bookmarkService: BookmarkService) {}

	@Post()
	async create(@Body() bookmarkDto: CreateBookmarkDto) {
		return await this.bookmarkService.createBookmark(bookmarkDto)
	}

	@Get()
	getAll() {
		return this.bookmarkService.getAllBookmarks()
	}
}
