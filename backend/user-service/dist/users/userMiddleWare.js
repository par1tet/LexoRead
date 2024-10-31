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
const jsonwebtoken_1 = require("jsonwebtoken");
let UserMiddleware = class UserMiddleware {
    use(req, res, next) {
        const secret = process.env.JWT_SECRET_KEY;
        console.log('JWT Secret:', secret);
        if (req.method === 'OPTIONS') {
            return next();
        }
        const token = jsonwebtoken_1.default.sign({ id: 1 }, secret);
        console.log('Generated Token:', token);
        const decoded = jsonwebtoken_1.default.verify(token, secret);
        console.log('Decoded Token:', decoded);
        if (!token) {
            return res.status(403).send({ msg: 'пользователь не авторизован' });
        }
        try {
            const decodedData = jsonwebtoken_1.default.verify(token, secret);
            console.log('Decoded Data:', decodedData);
            next();
        }
        catch (error) {
            console.error('JWT Verification Error:', error);
            return res.status(403).send({ msg: 'Invalid token' });
        }
    }
};
exports.UserMiddleware = UserMiddleware;
exports.UserMiddleware = UserMiddleware = __decorate([
    (0, common_1.Injectable)()
], UserMiddleware);
//# sourceMappingURL=userMiddleWare.js.map