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
  const [message, setMessage] = useState('')
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  const handleChange = (field: keyof FormState) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setForm((prev) => ({ ...prev, [field]: event.target.value }))
  }

  const handleLogin = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    if (!form.username || !form.password) {
      setStatus('error')
      setMessage('Username and password are required.')
      return
    }

    setStatus('loading')
    setMessage('Signing in...')
    try {
      const user = await login({ username: form.username, password: form.password })
      setStatus('ready')
      setMessage(`Signed in as ${user.name || user.username}.`)
      setIsLoggedIn(true)
      setForm({ username: '', password: '' })
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Login failed.'
      setStatus('error')
      setMessage(message)
    }
  }

  const handleSignup = async () => {
    if (!form.username || !form.password) {
      setStatus('error')
      setMessage('Username and password are required.')
      return
    }

    setStatus('loading')
    setMessage('Creating account...')
    try {
      await signup({ name: form.username, username: form.username, password: form.password })
      const user = await login({ username: form.username, password: form.password })
      setStatus('ready')
      setMessage(`Signed in as ${user.name || user.username}.`)
      setIsLoggedIn(true)
      setForm({ username: '', password: '' })
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Signup failed.'
      setStatus('error')
      setMessage(message)
    }
  }

  const handleLogout = async () => {
    setStatus('loading')
    setMessage('Signing out...')
    try {
      await logout()
      setStatus('idle')
      setMessage('Signed out.')
      setIsLoggedIn(false)
    } catch (error) {
      const message = error instanceof Error ? error.message : 'Logout failed.'
      setStatus('error')
      setMessage(message)
    }
  }

  return (
    <aside className="corner" aria-label="Authentication">
      <div className={`log-box status-${status}`} aria-live="polite">
        <p className="log-line">{message || 'Sign in to continue.'}</p>
      </div>
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
          {isLoggedIn ? (
            <button type="button" className="logout" onClick={handleLogout} disabled={status === 'loading'}>
              Logout
            </button>
          ) : null}
        </div>
      </form>
    </aside>
  )
}

export default SessionPanel
