<template>
  <UError class="-mt-10"
    :error="{
      statusCode: 404,
      statusMessage: 'Страница не найдена',
      message: 'Кажется, такой страницы не существует.'
    }"
    :clear="false"
  />
</template>
