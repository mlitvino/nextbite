import { useState } from 'react'
import { login, logout, signup } from '../services/auth'

type Status = 'idle' | 'loading' | 'ready' | 'error'

type FormState = {
  username: string
  password: string
}

function SessionPanel() {
  const [form, setForm] = useState<FormState>({ username: '', password: '' })
  const [status, setStatus] = useState<Status>('idle')
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [displayName, setDisplayName] = useState('')

  const handleChange = (field: keyof FormState) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setForm((prev) => ({ ...prev, [field]: event.target.value }))
  }

  const handleLogin = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    if (!form.username || !form.password) {
      setStatus('error')
      return
    }

    setStatus('loading')
    try {
      const user = await login({ username: form.username, password: form.password })
      setStatus('ready')
      setIsLoggedIn(true)
      setDisplayName(user.name || user.username)
      setForm({ username: '', password: '' })
    } catch (error) {
      setStatus('error')
    }
  }

  const handleSignup = async () => {
    if (!form.username || !form.password) {
      setStatus('error')
      return
    }

    setStatus('loading')
    try {
      await signup({ name: form.username, username: form.username, password: form.password })
      const user = await login({ username: form.username, password: form.password })
      setStatus('ready')
      setIsLoggedIn(true)
      setDisplayName(user.name || user.username)
      setForm({ username: '', password: '' })
    } catch (error) {
      setStatus('error')
    }
  }

  const handleLogout = async () => {
    setStatus('loading')
    try {
      await logout()
      setStatus('idle')
      setIsLoggedIn(false)
      setDisplayName('')
    } catch (error) {
      setStatus('error')
    }
  }

  return (
    <aside className="corner" aria-label="Authentication">
      {isLoggedIn ? (
        <div className="welcome">
          <p className="welcome-text">Welcome, {displayName}!</p>
          <button type="button" className="logout" onClick={handleLogout} disabled={status === 'loading'}>
            Logout
          </button>
        </div>
      ) : (
        <>
          <form className="login-form" onSubmit={handleLogin}>
            <input
              className="field-input"
              type="text"
              name="username"
              autoComplete="username"
              placeholder="Username"
              value={form.username}
              onChange={handleChange('username')}
            />
            <input
              className="field-input"
              type="password"
              name="password"
              autoComplete="current-password"
              placeholder="Password"
              value={form.password}
              onChange={handleChange('password')}
            />
            <div className="actions">
              <button type="submit" className="primary" disabled={status === 'loading'}>
                Login
              </button>
              <button type="button" className="secondary" onClick={handleSignup} disabled={status === 'loading'}>
                Signup
              </button>
            </div>
          </form>
        </>
      )}
    </aside>
  )
}

export default SessionPanel
