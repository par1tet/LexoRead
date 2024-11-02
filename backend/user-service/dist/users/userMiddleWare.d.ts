import { NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
import 'core-js/stable/atob';
declare global {
    namespace Express {
        interface Request {
            user?: any;
        }
    }
}
export declare class UserMiddleware implements NestMiddleware {
    use(req: Request, res: Response, next: NextFunction): void | Response<any, Record<string, any>>;
}
