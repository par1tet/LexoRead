import { ApiProperty } from "@nestjs/swagger";

export class changeDescriptionDto {
	@ApiProperty({example: 1})
	userId: number;
	@ApiProperty({example: 'Hello Man'})
	newDescription: string
}