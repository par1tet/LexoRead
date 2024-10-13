import { ApiProperty } from "@nestjs/swagger";

export class ChangeNameDto {
	@ApiProperty({example: 1})
	userId: number
	@ApiProperty({example: 'Maxim'})
	newUsername: string;
}