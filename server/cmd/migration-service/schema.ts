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

export const userRoles = pgEnum("user_roles", ["admin", "company", "user"]);
export const userInProjectRoles = pgEnum("user_in_project_roles", [
    "scrum_master",
    "developer",
    "qa",
    "ux",
    "ui",
    "pm",
]);

export const userLoginProviders = pgEnum("user_login_providers", [
    "email",
    "google",
    "github",
]);

export const companyTable = pgTable("companies", {
    id: uuid("id").primaryKey().defaultRandom(),
    name: varchar({ length: 255 }).notNull(),
});
export const usersTable = pgTable("users", {
    id: uuid("id").primaryKey().defaultRandom(),
    company_id: uuid("company_id").references(() => companyTable.id),
    first_name: varchar({ length: 255 }).notNull(),
    last_name: varchar({ length: 255 }).notNull(),
    email: varchar({ length: 255 }).notNull().unique(),
    role: userRoles("role").default("user").notNull(),
    password: varchar({ length: 255 }),
    is_active: boolean().default(false),
    is_premium: boolean().default(false),
    premium_start_date: date(),
    premium_end_date: date(),
    login_provider: userLoginProviders("login_provider").notNull(),
    otp_secret: varchar({ length: 10 }),
    created_at: timestamp().notNull().defaultNow(),
    updated_at: timestamp().notNull().defaultNow().$onUpdate(() => new Date()),
}, (table) => [
    uniqueIndex("email").on(table.email),
    uniqueIndex("company_id").on(table.company_id),
]);

export const projectTable = pgTable("projects", {
    id: uuid("id").primaryKey().defaultRandom(),
    company_id: uuid("company_id").references(() => companyTable.id),
    name: varchar({ length: 255 }).notNull(),
    description: varchar({ length: 255 }),
    created_at: timestamp().notNull().defaultNow(),
    updated_at: timestamp().notNull().defaultNow().$onUpdate(() => new Date()),
}, (table) => [
    uniqueIndex("company_id").on(table.company_id),
]);

export const projectMembersTable = pgTable("project_members", {
    id: uuid("id").primaryKey().defaultRandom(),
    project_id: uuid("project_id").references(() => projectTable.id),
    user_id: uuid("user_id").references(() => usersTable.id),
    role: userInProjectRoles("role").default("developer").notNull(),
    created_at: timestamp().notNull().defaultNow(),
    updated_at: timestamp().notNull().defaultNow().$onUpdate(() => new Date()),
}, (table) => [
    uniqueIndex("project_id").on(table.project_id),
    uniqueIndex("user_id").on(table.user_id),
    uniqueIndex("project_user").on(table.project_id, table.user_id),
]);
