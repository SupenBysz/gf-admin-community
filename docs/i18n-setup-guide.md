# 国际化(i18n)设置指南

本指南说明如何在前端项目中集成国际化(i18n)功能，包括语言切换组件和多语言支持。

## 后端API

后端已经提供了以下API接口用于语言切换：

1. `GET /api/common/getLanguages` - 获取支持的语言列表
2. `GET /api/common/getCurrentLanguage` - 获取当前使用的语言
3. `GET /api/common/setLanguage?lang=en-US` - 设置语言（支持的语言代码：`zh-CN`, `en-US`）

## 前端集成步骤

### 1. 安装依赖

如果您使用Vue.js + Ant Design Vue，请确保已安装以下依赖：

```bash
npm install vue axios ant-design-vue
# 或使用yarn
yarn add vue axios ant-design-vue
```

### 2. 添加语言切换组件

将提供的`LanguageSwitcher.vue`组件复制到您的项目中，例如放在`src/components/common/LanguageSwitcher.vue`。

### 3. 在布局中使用语言切换组件

在您的应用布局（如页面顶部导航栏）中添加语言切换组件：

```vue
<template>
  <div class="layout">
    <header>
      <div class="header-right">
        <!-- 其他头部内容 -->
        <language-switcher />
      </div>
    </header>
    <!-- 页面内容 -->
  </div>
</template>

<script>
import LanguageSwitcher from '@/components/common/LanguageSwitcher.vue'

export default {
  components: {
    LanguageSwitcher
  }
}
</script>
```

### 4. 设置国际化文案

前端项目中使用Vue-i18n管理国际化文案：

1. 安装Vue-i18n：

```bash
npm install vue-i18n@next
# 或使用yarn
yarn add vue-i18n@next
```

2. 创建语言文件：

在`src/locales`目录下创建以下文件：

**src/locales/zh-CN.js**:
```js
export default {
  common: {
    submit: '提交',
    cancel: '取消',
    confirm: '确认',
    delete: '删除',
    save: '保存',
    edit: '编辑',
    back: '返回',
    search: '搜索'
    // 更多通用文案...
  },
  login: {
    title: '登录',
    username: '用户名',
    password: '密码',
    remember: '记住我',
    forgot: '忘记密码?',
    submit: '登录',
    register: '注册账号'
    // 更多登录相关文案...
  },
  // 其他页面的文案...
}
```

**src/locales/en-US.js**:
```js
export default {
  common: {
    submit: 'Submit',
    cancel: 'Cancel',
    confirm: 'Confirm',
    delete: 'Delete',
    save: 'Save',
    edit: 'Edit',
    back: 'Back',
    search: 'Search'
    // More common phrases...
  },
  login: {
    title: 'Login',
    username: 'Username',
    password: 'Password',
    remember: 'Remember me',
    forgot: 'Forgot password?',
    submit: 'Login',
    register: 'Register'
    // More login related phrases...
  },
  // Other page texts...
}
```

3. 创建i18n实例并在main.js中注册：

**src/locales/index.js**:
```js
import { createI18n } from 'vue-i18n'
import zhCN from './zh-CN'
import enUS from './en-US'
import axios from 'axios'

const i18n = createI18n({
  legacy: false,
  locale: 'zh-CN', // 默认语言
  fallbackLocale: 'zh-CN', // 回退语言
  messages: {
    'zh-CN': zhCN,
    'en-US': enUS
  }
})

// 根据后端API返回的语言设置前端语言
export const setupI18n = async () => {
  try {
    const response = await axios.get('/api/common/getCurrentLanguage')
    if (response.data.code === 0) {
      i18n.global.locale.value = response.data.data
    }
  } catch (error) {
    console.error('获取当前语言失败:', error)
  }
  
  return i18n
}

export default i18n
```

**src/main.js**:
```js
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { setupI18n } from './locales'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'

const setupApp = async () => {
  const app = createApp(App)
  const i18n = await setupI18n()
  
  app.use(router)
  app.use(i18n)
  app.use(Antd)
  
  app.mount('#app')
}

setupApp()
```

### 5. 在组件中使用国际化文案

在Vue组件中使用`useI18n`钩子或`$t`函数获取国际化文案：

```vue
<template>
  <div>
    <h1>{{ t('login.title') }}</h1>
    <a-form>
      <a-form-item :label="t('login.username')">
        <a-input v-model:value="username" />
      </a-form-item>
      <a-form-item :label="t('login.password')">
        <a-input-password v-model:value="password" />
      </a-form-item>
      <a-form-item>
        <a-checkbox v-model:checked="remember">{{ t('login.remember') }}</a-checkbox>
        <a href="#">{{ t('login.forgot') }}</a>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit">{{ t('login.submit') }}</a-button>
        <a-button style="margin-left: 10px">{{ t('common.cancel') }}</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

export default {
  setup() {
    const { t } = useI18n()
    const username = ref('')
    const password = ref('')
    const remember = ref(false)
    
    return {
      t,
      username,
      password,
      remember
    }
  }
}
</script>
```

## 注意事项

1. 确保API请求中正确处理语言切换后的刷新，以应用新的语言设置。
2. 在开发过程中维护好语言文件，确保所有文案都有对应的翻译。
3. 可以考虑使用自动化工具提取和管理翻译文案，减少手动维护的工作量。
4. 前端和后端的语言代码需要一致，例如都使用`zh-CN`和`en-US`作为语言标识。

## 扩展功能

1. 添加更多语言支持：在后端`SysCommonController.GetLanguages`方法和前端语言文件中添加更多语言。
2. 记住用户语言偏好：可以将用户选择的语言存储在后端用户配置中，以便在不同设备上保持一致的语言设置。
3. 基于用户地理位置自动选择语言：可以使用IP地址解析服务确定用户可能的地理位置，并据此提供默认语言设置。 