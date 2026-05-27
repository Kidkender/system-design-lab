import { apiClient } from '@/lib/api/client'
import type { User } from '@/types/api'

export async function createUser(username: string): Promise<User> {
  const email = `${username.toLowerCase().replace(/\s+/g, '_')}@player.local`
  const res = await apiClient.post<User>('/users', {
    name: username,
    email,
    password: 'pixel-dungeon',
  })
  return res.data
}
