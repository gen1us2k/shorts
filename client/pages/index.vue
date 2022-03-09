<template>
  <!-- Container -->
  <div class="w-full min-h-screen bg-gray-100">
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
            v-model="url"
          />
        </div>
        <button
          v-if="!loading"
          @click="shortenUrl"
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

      <!-- Shortened Url -->
      <div class="py-6 w-full">
        <div
          class="my-3 bg-white shadow rounded-lg"
          v-for="(shortenedUrl, index) in shortenedUrls"
          :key="index"
        >
          <div class="px-4 py-5 sm:p-6">
            <div class="flex items-center">
              <svg
                class="h-5 w-5 text-green-600"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"
                />
              </svg>
              <h3 class=" px-2 text-lg leading-6 font-semibold text-gray-900">
                Your Shortened Url
              </h3>
            </div>
            <div
              class="mt-2 max-w-xl text-xs font-medium leading-5 text-gray-500 overflow-hidden"
            >
              <p>
                {{ shortenedUrl.url }}
              </p>
            </div>

            <div class="mt-3 text-sm leading-5">
              <a
                v-bind:href="shortenedUrl.shortened"
                target="_blank"
                class="font-medium text-indigo-600 hover:text-indigo-500 transition ease-in-out duration-150"
              >
                {{ shortenedUrl.shortened }} &rarr;
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage'
})
</script>
