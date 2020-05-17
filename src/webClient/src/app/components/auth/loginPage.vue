<template>
  <div class="layout">
    <div>
      <div class="row justify-center" style="margin-top: 20px; margin-bottom: 10px">
        <img
          src="https://cdn.pixabay.com/photo/2017/05/05/00/15/kokopelli-2285538_960_720.png"
          alt="" style="width: auto; max-height: 100px">
      </div>
      <!--Кнопки авторизации-->
      <div class="row justify-center q-mt-lg">
        <email-btn/>
      </div>
    </div>

  </div>
</template>

<script>
    import emailBtn from './email/emailAuthBtn'

    export default {
        components: {emailBtn},
        data() {
            return {}
        },
    }
</script>
