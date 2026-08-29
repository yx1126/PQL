<script setup lang="ts">
import ErrGif from "@/assets/image/403.gif";
defineOptions({
    name: "403",
});
const router = useRouter();
const { t } = useLocal();
const errGif = ErrGif + "?" + +new Date();

function back() {
    router.back();
}
</script>

<template>
    <el-card shadow="never" class="h-full">
        <div class="four-zero-three">
            <el-button type="primary" icon="EleArrowLeft" @click="back">{{ $t("button.back") }}</el-button>
            <el-row>
                <el-col :span="12">
                    <h1 class="text-jumbo text-ginormous">Oops!</h1>
                    <h2 class="text-h2">{{ t("access") }}</h2>
                    <h6 class="text-h6">{{ t("desc") }}</h6>
                    <ul class="list-unstyled">
                        <li class="link-type">
                            <el-button link type="primary" @click="back">{{ t("back") }}</el-button>
                        </li>
                    </ul>
                </el-col>
                <el-col :span="12">
                    <img :src="errGif" width="313" height="428" alt="Girl has dropped her ice cream.">
                </el-col>
            </el-row>
        </div>
    </el-card>
</template>

<style lang="scss" scoped>
.four-zero-three {
    width: 800px;
    max-width: 100%;
    margin: 100px auto;
    .pan-gif {
        margin: 0 auto;
        display: block;
    }
    .pan-img {
        display: block;
        margin: 0 auto;
        width: 100%;
    }
    .text-jumbo {
        font-size: 60px;
        font-weight: 700;
        color: #484848;
        margin: 40px 0;
    }
    .text-h2 {
        margin: 15px 0;
    }
    .text-h6 {
        margin: 20px 0;
    }
}
</style>

<i18n lang="yaml">
zh:
  access: 您没有访问权限！
  desc: 对不起，您没有访问权限，请不要进行非法操作！您可以返回主页面
  back: 回首页
en:
  access: You do not have access permissions!
  desc: Sorry, you do not have access rights, please do not perform illegal operations! You can return to the main page
  back: Back to Home
</i18n>