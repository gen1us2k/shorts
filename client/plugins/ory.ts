import { Configuration, V0alpha2Api } from '@ory/kratos-client'
import type { Session, SelfServiceLogoutUrl } from '@ory/kratos-client'

export default ({ app }, inject) => {
  const ory = new V0alpha2Api(
    new Configuration({
      basePath: app.$config.kratosAPIURL,
      baseOptions: {
        withCredentials: true
      }
    })
  )
  ory.toSession()
    .then(({ data }) => {
      inject('session', data)
      inject('authenticated', true)
    })
    .catch(() => {
      inject('authenticated', false)
      console.log('Unauthenticated')
    })
  inject('ory', {
    authenticated: true
  })
}
