package org.example;

import java.nio.charset.StandardCharsets;
import java.util.Base64;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello world!");
    }


    public CodeInfoBo parseQidianAuthCode(String code) {
        if (StringUtils.isBlank(code)) {
            return null;
        }
        String raw = new String(Base64.getDecoder().decode(code.trim()), StandardCharsets.UTF_8);
        String[] split = raw.split("_");
        if (split.length < 6) {
            return null;
        }

        try {
            CodeInfoBo codeInfo = new CodeInfoBo();
            System.out.println("", );
            codeInfo.setTimestampMillis(Long.parseLong(split[0]));
            codeInfo.setConsultSource(Integer.parseInt(split[1]));
            codeInfo.setToken(split[2]);
            codeInfo.setUserId(Long.parseLong(split[3]));
            codeInfo.setUid(Long.parseLong(split[4]));
            codeInfo.setUrl(parsePathCode(split[5], code));

            return codeInfo;
        } catch (Exception e) {
            return null;
        }
    }

     class CodeInfoBo {

//        @ApiModelProperty("时间戳，毫秒")
        private Long timestampMillis;

//        @ApiModelProperty("当前登录用户ID")
        private Long userId;
//        @ApiModelProperty("当前登录账户ID")
        private Long uid;

//        @ApiModelProperty("当前会话来源的平台：0-未知，1-投放端PC，2-投放端小程序，3-官网")
        private Integer consultSource;

//        @ApiModelProperty("当前会话来源的平台登录态标，如果是投放端或者服务商系统PC页面，则是登录后cookie里的gdt_token，如果是公众号或者小程序则是appid|openid等，没有登录或者没有标识填0")
        private String token;

//        @ApiModelProperty("当前反馈页面的uri path，注意需要先base64后再用“_”拼接，防止特殊字符，如果取不到path，则对0进行base64后再用“_”拼接")
        private String url;

    }
}