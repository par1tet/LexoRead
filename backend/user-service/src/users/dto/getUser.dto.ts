import { ApiProperty } from "@nestjs/swagger";

export class GetUserDto {
	@ApiProperty({example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5cHax5RecXb3jpX35uUS9wT50MLI5zVTyz9Ludm2AxvlvyDA5kt6AjVzasSoAXMLm7'})
	jwtToken: string
}