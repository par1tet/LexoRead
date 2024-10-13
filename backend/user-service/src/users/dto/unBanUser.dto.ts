import { ApiProperty } from "@nestjs/swagger";

export class UnBanUserDto {
	@ApiProperty({example: 1})
	id: number;
}