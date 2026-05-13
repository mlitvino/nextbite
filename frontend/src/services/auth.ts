import type { LoginPayload, SignupPayload, User } from '../types/auth'
import { post } from './api'

export async function login(payload: LoginPayload): Promise<User> {
  return post<User>('/auth/login', payload)
}

export async function logout(): Promise<void> {
  await post<void>('/auth/logout')
}

export async function signup(payload: SignupPayload): Promise<User> {
  return post<User>('/auth/signup', payload)
}
