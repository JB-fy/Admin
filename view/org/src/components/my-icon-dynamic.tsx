import * as epIconList from '@element-plus/icons-vue'

const iconMap: { [propName: string]: any } = {
    // 'autoicon-mdi-account': <autoicon-mdi-account />,
}
for (const [key, component] of Object.entries(epIconList)) {
    iconMap['autoicon-ep-' + key.replace(/([a-z])([A-Z])/g, '$1-$2').toLowerCase()] = component
}

export default iconMap
