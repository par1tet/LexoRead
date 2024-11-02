"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserMiddleware = void 0;
const common_1 = require("@nestjs/common");
const jwt = require("jsonwebtoken");
require("core-js/stable/atob");
let UserMiddleware = class UserMiddleware {
    use(req, res, next) {
        try {
            const secret = process.env.JWT_SECRET_KEY;
            if (req.method === 'OPTIONS') {
                return next();
            }
            const authHeader = req.headers.authorization;
            if (!authHeader) {
                return res.status(403).send({ msg: 'пользователь не авторизован' });
            }
            const token = authHeader.split(' ')[1];
            const decodedData = jwt.verify(token, secret, { algorithms: ['HS256'] });
            next();
            req.user = decodedData;
        }
        catch (err) {
            console.log(err);
        }
    }
};
exports.UserMiddleware = UserMiddleware;
exports.UserMiddleware = UserMiddleware = __decorate([
    (0, common_1.Injectable)()
], UserMiddleware);
//# sourceMappingURL=userMiddleWare.js.map