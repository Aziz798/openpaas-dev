import {
    boolean,
    date,
    pgEnum,
    pgTable,
    timestamp,
    uniqueIndex,
    uuid,
    varchar,
} from "drizzle-orm/pg-core";
export const userRoles = pgEnum("user_roles", ["admin","company","user"]);

export const usersTable = pgTable("users", {
    id: uuid("id").primaryKey().defaultRandom(),
    first_name: varchar({ length: 255 }).notNull(),
    last_name: varchar({ length: 255 }).notNull(),
    email: varchar({ length: 255 }).notNull().unique(),
    role: userRoles("role").default("user").notNull(),
    password: varchar({ length: 255 }),
    is_active: boolean().default(false),
    is_premium: boolean().default(false),
    premium_start_date: date(),
    premium_end_date: date(),
    login_provider: varchar({ length: 15 }).notNull(),
    otp_secret: varchar({ length: 10 }),
    created_at: timestamp().notNull().defaultNow(),
    updated_at: timestamp().notNull().defaultNow().$onUpdate(() => new Date()),
}, (table) => [
    uniqueIndex("email").on(table.email),
]);
