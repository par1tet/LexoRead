import { ApiProperty } from "@nestjs/swagger";

export class changeAvatarUrlDto {
	@ApiProperty({example: 1})
	id: number;
	@ApiProperty({example: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQy1QKMwvvFRtE7kNDp0oyv8e5k39uBdPxbbg&s'})
	newAvatarUrl: string
}