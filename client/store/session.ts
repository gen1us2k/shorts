export const state = () => ({
  session: {},
  authenticated: false
})

export const mutations = {
  setSession (state, session) {
    state.session = session
  },
  setAuthenticated (state, authenticated) {
    state.authenticated = authenticated
  }
}
