frontProjectPath="/Users/mtgnorton/Coding/vue/gf-admin-ui"
backendProjectPath="/Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin"
packName="binance-collect"

#echo "${backendProjectPath}/public/front/*"
# rm -rf   ${backendProjectPath}/public/front/*

#
# rm -rf /Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin/public/front/*


scp -r $backendProjectPath/bin/linux_amd64/$packName 261:/www/wwwroot/binance-collect
