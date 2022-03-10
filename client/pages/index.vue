<template>
  <!-- Container -->
  <div class="w-full min-h-screen bg-gray-100">
  <AppHeader :session="$data.session" :authenticated="authenticated" />
    <div
      class="max-w-2xl mx-auto min-h-screen flex flex-col items-center justify-center px-4"
    >
      <!-- Logo Image -->
      <ShortsLogo />
      <!-- Header -->
      <h1 class="text-gray-900 font-black text-5xl uppercase text-center">
        Shorts
      </h1>
      <h2 class="text-indigo-700 text-sm font-semibold italic">
        Make your URLs shorter!
      </h2>

      <!-- Url Input -->
      <div class=" w-full mt-6 flex rounded-md shadow-sm">
        <div class="relative flex-grow focus-within:z-10">
          <DownloadIcon />
          <input
            class="text-gray-700 py-3 form-input block w-full rounded-none rounded-l-md pl-10 transition ease-in-out duration-150 font-semibold sm:text-sm sm:leading-5"
            placeholder="www.example.com"
          />
        </div>
        <button
          v-if="!loading"
          class="group -ml-px relative inline-flex items-center px-3 py-3 border border-indigo-300 text-sm leading-5 font-medium rounded-r-md text-white bg-indigo-700 hover:text-indigo-700 hover:bg-white focus:outline-none focus:shadow-outline-blue focus:border-indigo-300 active:bg-gray-100 active:text-indigo-700 transition ease-in-out duration-150"
        >
          <svg
            class="text-white h-5 w-5 group-hover:text-indigo-700"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 13l-3 3m0 0l-3-3m3 3V8m0 13a9 9 0 110-18 9 9 0 010 18z"
            />
          </svg>
          <span class="ml-2 text-sm font-semibold">Shorten</span>
        </button>
        <button
          v-else
          class="group -ml-px relative inline-flex items-center px-3 py-3 border border-indigo-300 text-sm leading-5 font-medium rounded-r-md text-white bg-indigo-700 hover:text-indigo-700 hover:bg-white focus:outline-none focus:shadow-outline-blue focus:border-indigo-300 active:bg-gray-100 active:text-indigo-700 transition ease-in-out duration-150"
        >
          <svg
            class="text-white h-5 w-5 group-hover:text-indigo-700 animate-spin"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2-1M4 7l2 1M4 7v2.5M12 21l-2-1m2 1l2-1m-2 1v-2.5M6 18l-2-1v-2.5M18 18l2-1v-2.5"
            />
          </svg>
          <span class="ml-2 text-sm font-semibold">Shortening</span>
        </button>
      </div>

      <!-- Error Message -->
      <p
        v-show="errorMessage"
        class="text-xs font-semibold text-red-600 italic"
      >
        {{ errorMessage }}
      </p>

    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Configuration, V0alpha2Api } from '@ory/kratos-client'
import type { SelfServiceLogoutUrl } from '@ory/kratos-client'
import type { AxiosResponse } from 'axios'
import AppHeader from '../components/AppHeader'

const getAuthState = async ({ app }) => {
  const ory = new V0alpha2Api(
    new Configuration({
      basePath: app.$config.kratosAPIURL,
      baseOptions: {
        withCredentials: true
      }
    })
  )

  try {
    const session = await ory.toSession()
    return {
      authenticated: true,
      session: session,
    }
  } catch {
    return {
      authenticated: false,
      session: {},
    }
  }
}

export default Vue.extend({
  name: 'IndexPage',
  data: () => {
    return {
      authenticated: false,
      session: {},
      loading: false,
      errorMessage: ""
    }
  },
  async asyncData (context) {
		const authState = await getAuthState(context)
    return {
      ...authState
    }
  },
})
</script>
